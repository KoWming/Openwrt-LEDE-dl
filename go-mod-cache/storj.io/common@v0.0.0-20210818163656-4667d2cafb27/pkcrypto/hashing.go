// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package pkcrypto

import (
	"crypto/sha256"
	"hash"
)

// NewHash returns default hash in storj.
func NewHash() hash.Hash {
	return sha256.New()
}

// SHA256Hash calculates the SHA256 hash of the input data.
func SHA256Hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
