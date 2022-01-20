// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package rpcstatus

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/drpc/drpcerr"
)

var allCodes = []StatusCode{
	Unknown,
	OK,
	Canceled,
	InvalidArgument,
	DeadlineExceeded,
	NotFound,
	AlreadyExists,
	PermissionDenied,
	ResourceExhausted,
	FailedPrecondition,
	Aborted,
	OutOfRange,
	Unimplemented,
	Internal,
	Unavailable,
	DataLoss,
	Unauthenticated,
}

func TestStatus(t *testing.T) {
	for _, code := range allCodes {
		err := Error(code, "")
		assert.Equal(t, Code(err), code)
		assert.Equal(t, drpcerr.Code(err), uint64(code))
	}

	assert.Equal(t, Code(nil), OK)
	assert.Equal(t, Code(context.Canceled), Canceled)
	assert.Equal(t, Code(context.DeadlineExceeded), DeadlineExceeded)
}

func TestStatus_WrapFormatting(t *testing.T) {
	err := Wrap(Internal, errors.New("test"))
	assert.True(t, strings.Count(fmt.Sprintf("%+v", err), "\n") > 0)
}
