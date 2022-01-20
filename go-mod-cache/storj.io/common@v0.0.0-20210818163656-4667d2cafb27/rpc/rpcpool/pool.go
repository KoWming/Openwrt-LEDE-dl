// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package rpcpool

import (
	"context"
	"crypto/tls"
	"runtime"
	"time"

	"storj.io/common/peertls/tlsopts"
	"storj.io/common/rpc/rpccache"
	"storj.io/drpc"
)

// Options controls the options for a connection pool.
type Options struct {
	// Capacity is how many connections to keep open.
	Capacity int

	// KeyCapacity is the number of connections to keep open per cache key.
	KeyCapacity int

	// IdleExpiration is how long a connection in the pool is allowed to be
	// kept idle. If zero, connections do not expire.
	IdleExpiration time.Duration
}

// Pool is a wrapper around a cache of connections that allows one to get or
// create new cached connections.
type Pool struct {
	cache *rpccache.Cache
}

// New constructs a new Pool with the Options.
func New(opts Options) *Pool {
	p := &Pool{cache: rpccache.New(rpccache.Options{
		Expiration:  opts.IdleExpiration,
		Capacity:    opts.Capacity,
		KeyCapacity: opts.KeyCapacity,
		Close: func(pv interface{}) error {
			return pv.(*poolValue).conn.Close()
		},
		Stale: func(pv interface{}) bool {
			select {
			case <-pv.(*poolValue).conn.Closed():
				return true
			default:
				return false
			}
		},
	})}

	// As much as I dislike finalizers, especially for cases where it handles
	// file descriptors, I think it's important to add one here at least until
	// a full audit of all of the uses of the rpc.Dialer type and ensuring they
	// all get closed.
	runtime.SetFinalizer(p, func(p *Pool) {
		mon.Event("pool_leaked")
		_ = p.Close()
	})

	return p
}

// poolKey is the type of keys in the cache.
type poolKey struct {
	key        string
	tlsOptions *tlsopts.Options
}

// poolValue is the type of values in the cache.
type poolValue struct {
	conn  drpc.Conn
	state *tls.ConnectionState
}

// Dialer is the type of function to create a new connection.
type Dialer = func(context.Context) (drpc.Conn, *tls.ConnectionState, error)

// Close closes all of the cached connections. It is safe to call on a nil receiver.
func (p *Pool) Close() error {
	if p == nil {
		return nil
	}

	runtime.SetFinalizer(p, nil)
	return p.cache.Close()
}

// get returns a drpc connection from the cache if possible, dialing if necessary.
func (p *Pool) get(ctx context.Context, pk poolKey, dial Dialer) (pv *poolValue, err error) {
	defer mon.Task()(&ctx)(&err)

	if p != nil {
		pv, ok := p.cache.Take(pk).(*poolValue)
		if ok {
			mon.Event("connection_from_cache")
			return pv, nil
		}
	}

	mon.Event("connection_dialed")
	conn, state, err := dial(ctx)
	if err != nil {
		return nil, err
	}

	return &poolValue{
		conn:  conn,
		state: state,
	}, nil
}

// Get looks up a connection with the same key and TLS options and returns it if it
// exists. If it does not exist, it calls the dial function to create one. It is safe
// to call on a nil receiver, and if so, always returns a dialed connection.
func (p *Pool) Get(ctx context.Context, key string, tlsOptions *tlsopts.Options, dial Dialer) (
	conn drpc.Conn, state *tls.ConnectionState, err error) {
	defer mon.Task()(&ctx)(&err)

	pk := poolKey{
		key:        key,
		tlsOptions: tlsOptions,
	}

	pv, err := p.get(ctx, pk, dial)
	if err != nil {
		return nil, nil, err
	}

	// if we have a nil pool, we always dial once and do not return a wrapped connection.
	if p == nil {
		return pv.conn, pv.state, nil
	}

	// we immediately place the connection back into the pool so that it may be used
	// by the returned poolConn.
	p.cache.Put(pk, pv)

	return &poolConn{
		ch:   make(chan struct{}),
		pk:   pk,
		dial: dial,
		pool: p,
	}, pv.state, nil
}
