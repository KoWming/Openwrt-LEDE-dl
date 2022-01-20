// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

// Package testrand implements generating random base types for testing.
package testrand

import (
	"io"
	"math/rand"

	"storj.io/common/memory"
	"storj.io/common/storj"
	"storj.io/common/uuid"
)

// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n <= 0.
func Intn(n int) int { return rand.Intn(n) }

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It panics if n <= 0.
func Int63n(n int64) int64 {
	return rand.Int63n(n)
}

// Float64n returns floating point pseudo-random number in [-n,0] || [0,n]
// based on the sign of the input.
func Float64n(n int64) float64 {
	return rand.Float64() * float64(n)
}

// Read reads pseudo-random data into data.
func Read(data []byte) {
	const newSourceThreshold = 64
	if len(data) < newSourceThreshold {
		/* #nosec G404 */ // This package is only used for testing
		_, _ = rand.Read(data)
		return
	}

	src := rand.NewSource(rand.Int63())
	r := rand.New(src)
	_, _ = r.Read(data)
}

// Bytes generates size amount of random data.
func Bytes(size memory.Size) []byte {
	data := make([]byte, size.Int())
	Read(data)
	return data
}

// BytesInt generates size amount of random data.
func BytesInt(size int) []byte {
	return Bytes(memory.Size(size))
}

// Reader creates a new random data reader.
func Reader() io.Reader {
	return rand.New(rand.NewSource(rand.Int63()))
}

// NodeID creates a random node id.
func NodeID() storj.NodeID {
	var id storj.NodeID
	Read(id[:])
	// set version to 0
	id[len(id)-1] = 0
	return id
}

// PieceID creates a random piece id.
func PieceID() storj.PieceID {
	var id storj.PieceID
	Read(id[:])
	return id
}

// Key creates a random test key.
func Key() storj.Key {
	var key storj.Key
	Read(key[:])
	return key
}

// Nonce creates a random test nonce.
func Nonce() storj.Nonce {
	var nonce storj.Nonce
	Read(nonce[:])
	return nonce
}

// SerialNumber creates a random serial number.
func SerialNumber() storj.SerialNumber {
	var serial storj.SerialNumber
	Read(serial[:])
	return serial
}

// StreamID creates a random stream ID.
func StreamID(size int) storj.StreamID {
	return storj.StreamID(BytesInt(size))
}

// SegmentID creates a random segment ID.
func SegmentID(size int) storj.SegmentID {
	return storj.SegmentID(BytesInt(size))
}

// UUID creates a random uuid.
func UUID() uuid.UUID {
	var uuid uuid.UUID
	Read(uuid[:])
	return uuid
}

// BucketName creates a random bucket name mostly confirming to the
// restrictions of S3:
// https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html
//
// NOTE: This may not generate values that cover all valid values (for Storj or
// S3). This is a best effort to cover most cases we believe our design
// requires and will need to be revisited when a more explicit design spec is
// created.
func BucketName() string {
	const (
		edges = "abcdefghijklmnopqrstuvwxyz0123456789"
		body  = "abcdefghijklmnopqrstuvwxyz0123456789-"
		min   = 3
		max   = 63
	)

	size := rand.Intn(max-min) + min

	b := make([]byte, size)
	for i := range b {
		switch i {
		case 0:
			fallthrough
		case size - 1:
			b[i] = edges[rand.Intn(len(edges))]
		default:
			b[i] = body[rand.Intn(len(body))]
		}
	}

	return string(b)
}

// Metadata creates random metadata mostly conforming to the restrictions of S3:
// https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-metadata
//
// NOTE: This may not generate values that cover all valid values (for Storj or
// S3). This is a best effort to cover most cases we believe our design
// requires and will need to be revisited when a more explicit design spec is
// created.
func Metadata() map[string]string {
	const (
		// The actual limit is 2KiB, but there are overheads to encoding and encryption.
		max = 1 * 1024
	)

	total := rand.Intn(max)
	metadata := make(map[string]string)

	for used := 0; total-used > 1; {
		keySize := rand.Intn(total-(used+1)) + 1
		key := BytesInt(keySize)
		used += len(key)

		valueSize := rand.Intn(total - used)
		value := BytesInt(valueSize)
		used += len(value)

		metadata[string(key)] = string(value)
	}

	return metadata
}

const (
	letters    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers    = "0123456789"
	safepath   = "-_."
	unsafepath = "[]()^ #%&!@:+={}'~"
	// things that are non-printable and isn't special for path component encoding.
	nonprintable = "\x03\xfd\xf0\x2d\x30"
)

// URLPath creates a clean random url path.
func URLPath() string {
	n := rand.Intn(1024) + 1
	path := bytesOf(n, letters+numbers+numbers+safepath+safepath+safepath)

	// replace / every 3 to 100 characters
	offset := randomRange(3, 100)
	for offset < len(path) {
		path[offset] = '/'
		offset += randomRange(3, 100)
	}

	return string(path)
}

// URLPathNonFolder creates a clean random url path that is not a folder.
func URLPathNonFolder() string {
	n := rand.Intn(1024) + 1
	path := bytesOf(n, letters+numbers+numbers+safepath+safepath+safepath)

	// replace / every 3 to 100 characters
	offset := randomRange(3, 100)
	for offset < len(path)-1 {
		path[offset] = '/'
		offset += randomRange(3, 100)
	}

	return string(path)
}

// Path creates a random path that Storj should support.
func Path() string {
	n := rand.Intn(1024) + 1
	path := bytesOf(n, letters+numbers+numbers+safepath+safepath+safepath+unsafepath+"/////"+nonprintable+nonprintable+nonprintable)
	return string(path)
}

// bytesOf generates random bytes using the specified alphabet.
func bytesOf(length int, alphabet string) []byte {
	if len(alphabet) > 256 {
		panic("alphabet too large")
	}

	data := BytesInt(length)
	for i, v := range data {
		data[i] = alphabet[int(v)%len(alphabet)]
	}
	return data
}

func randomRange(min, max int) int {
	return rand.Intn(max-min) + min
}
