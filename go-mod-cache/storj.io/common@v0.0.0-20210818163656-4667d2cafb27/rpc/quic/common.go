// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

// +build go1.15

package quic

import (
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"

	"storj.io/common/rpc"
)

var (
	mon = monkit.Package()

	// Error is a pkg/quic error.
	Error = errs.Class("quic")
)

const quicConnectorPriority = 20

func init() {
	rpc.RegisterCandidateConnectorType("quic", func() rpc.Connector {
		return NewDefaultConnector(nil)
	}, quicConnectorPriority)
}
