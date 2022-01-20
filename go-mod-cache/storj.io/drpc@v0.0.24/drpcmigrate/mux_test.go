// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpcmigrate

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/zeebo/assert"
	"github.com/zeebo/errs"
)

func TestMux(t *testing.T) {
	run := func(lis net.Listener, data string) error {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}

		buf := make([]byte, len(data))
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			return err
		}

		assert.Equal(t, data, string(buf))
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis := newFakeListener(
		newPrefixConn([]byte("prefix1data1"), nil),
		newPrefixConn([]byte("prefix2data2"), nil),
		newPrefixConn([]byte("prefix3data3"), nil),
	)

	mux := NewListenMux(lis, len("prefixN"))

	lisErrs := make(chan error, 3)
	go func() { lisErrs <- run(mux.Route("prefix1"), "data1") }()
	go func() { lisErrs <- run(mux.Route("prefix2"), "data2") }()
	go func() { lisErrs <- run(mux.Default(), "prefix3data3") }()

	muxErrs := make(chan error, 1)
	go func() { muxErrs <- mux.Run(ctx) }()

	for i := 0; i < 3; i++ {
		assert.NoError(t, <-lisErrs)
	}

	cancel()

	for i := 0; i < 1; i++ {
		assert.NoError(t, <-muxErrs)
	}
}

func TestMuxAcceptError(t *testing.T) {
	err := errs.New("problem")
	mux := NewListenMux(newErrorListener(err), 0)
	assert.Equal(t, mux.Run(context.Background()), err)
}

//
// fake listener
//

type fakeListener struct {
	done  chan struct{}
	err   error
	conns []net.Conn
}

func (fl *fakeListener) Addr() net.Addr { return nil }

func (fl *fakeListener) Close() error {
	close(fl.done)
	return nil
}

func (fl *fakeListener) Accept() (c net.Conn, err error) {
	if fl.err != nil {
		return nil, fl.err
	}
	if len(fl.conns) == 0 {
		<-fl.done
		return nil, Closed
	}
	c, fl.conns = fl.conns[0], fl.conns[1:]
	return c, nil
}

func newFakeListener(conns ...net.Conn) *fakeListener {
	return &fakeListener{
		done:  make(chan struct{}),
		conns: conns,
	}
}

func newErrorListener(err error) *fakeListener {
	return &fakeListener{err: err}
}
