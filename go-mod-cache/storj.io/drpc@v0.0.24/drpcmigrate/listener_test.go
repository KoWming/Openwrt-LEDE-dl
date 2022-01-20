// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package drpcmigrate

import (
	"net"
	"testing"

	"github.com/zeebo/assert"
)

func TestListener(t *testing.T) {
	type addr struct{ net.Addr }
	type conn struct{ net.Conn }

	lis := newListener(addr{})

	{ // ensure the addr is the same we passed in
		assert.Equal(t, lis.Addr(), addr{})
	}

	{ // ensure that we can accept a connection from the listener
		go func() { lis.Conns() <- conn{} }()
		c, err := lis.Accept()
		assert.NoError(t, err)
		assert.DeepEqual(t, c, conn{})
	}

	{ // ensure that closing the listener is no problem
		assert.NoError(t, lis.Close())
	}

	{ // ensure that accept after close returns the right error
		c, err := lis.Accept()
		assert.Equal(t, err, Closed)
		assert.Nil(t, c)
	}
}
