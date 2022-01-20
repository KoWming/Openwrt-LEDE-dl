// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package ranger

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"storj.io/common/testcontext"
)

func TestFileRanger(t *testing.T) {
	ctx := testcontext.New(t)

	for i, example := range []struct {
		data                 string
		size, offset, length int64
		substr               string
		fail                 bool
	}{
		{"", 0, 0, 0, "", false},
		{"abcdef", 6, 0, 0, "", false},
		{"abcdef", 6, 3, 0, "", false},
		{"abcdef", 6, 0, 6, "abcdef", false},
		{"abcdef", 6, 0, 5, "abcde", false},
		{"abcdef", 6, 0, 4, "abcd", false},
		{"abcdef", 6, 1, 4, "bcde", false},
		{"abcdef", 6, 2, 4, "cdef", false},
		{"abcdefg", 7, 1, 4, "bcde", false},
		{"abcdef", 6, 0, 7, "", true},
		{"abcdef", 6, -1, 7, "abcde", true},
		{"abcdef", 6, 0, -1, "abcde", true},
	} {
		fh, err := os.Create(ctx.File(strconv.Itoa(i)))
		if err != nil {
			t.Fatalf("failed making tempfile")
		}
		_, err = fh.Write([]byte(example.data))
		if err != nil {
			t.Fatalf("failed writing data")
		}
		name := fh.Name()
		err = fh.Close()
		if err != nil {
			t.Fatalf("failed closing data")
		}

		rr, err := FileRanger(name)
		if err != nil {
			t.Fatalf("failed opening tempfile")
		}
		if rr.Size() != example.size {
			t.Fatalf("invalid size: %v != %v", rr.Size(), example.size)
		}

		r, err := rr.Range(context.Background(), example.offset, example.length)
		if example.fail {
			if err == nil {
				t.Fatalf("expected error")
			}
			return
		}

		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}

		data, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(data, []byte(example.substr)) {
			t.Fatalf("invalid subrange: %#v != %#v", string(data), example.substr)
		}

		if err := r.Close(); err != nil {
			t.Fatalf("unable to close file %q: %v", name, err)
		}
	}
}

func TestFileRangerOpenFileError(t *testing.T) {
	rr, err := FileRanger("")
	if rr != nil {
		t.Fatal("Ranger expected to be nil")
	}
	if err == nil {
		t.Fatal("Error expected")
	}
}
