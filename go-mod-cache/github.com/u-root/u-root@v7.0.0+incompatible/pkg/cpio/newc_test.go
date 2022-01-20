// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpio

import (
	"bytes"
	"io"
	"io/ioutil"
	"reflect"
	"syscall"
	"testing"

	"github.com/u-root/u-root/pkg/uio"
)

/*
TODO(#935): This test should really compare against []Records, not the string
formatted []Records.

func TestSimple(t *testing.T) {
	r := Newc.Reader(bytes.NewReader(testCPIO))
	files, err := ReadAllRecords(r)
	if err != nil {
		t.Fatal(err)
	}

	for i, f := range files {
		if f.String() != testResult[i] {
			t.Errorf("Value %d: got \n%s, want \n%s", i, f.String(), testResult[i])
		}
		t.Logf("Value %d: got \n%s, want \n%s", i, f.String(), testResult[i])
	}
}
*/

func TestWriteRead(t *testing.T) {
	contents := []byte("LANAAAAAAAAAA")
	rec := StaticRecord(contents, Info{
		Ino:      1,
		Mode:     syscall.S_IFREG | 2,
		UID:      3,
		GID:      4,
		NLink:    5,
		MTime:    6,
		FileSize: 7,
		Major:    8,
		Minor:    9,
		Rmajor:   10,
		Rminor:   11,
		Name:     "foobar",
	})

	buf := &bytes.Buffer{}
	w := Newc.Writer(buf)
	if err := w.WriteRecord(rec); err != nil {
		t.Errorf("Could not write record %q: %v", rec.Name, err)
	}

	if err := WriteTrailer(w); err != nil {
		t.Errorf("Could not write trailer: %v", err)
	}

	r := Newc.Reader(bytes.NewReader(buf.Bytes()))
	rec2, err := r.ReadRecord()
	if err != nil {
		t.Errorf("Could not read record: %v", err)
	}

	if rec2.Info != rec.Info {
		t.Errorf("Records not equal:\n%#v\n%#v", rec.Info, rec2.Info)
	}

	contents2, err := ioutil.ReadAll(uio.Reader(rec2))
	if err != nil {
		t.Errorf("Could not read %q: %v", rec2.Name, err)
	}

	if !bytes.Equal(contents2, contents) {
		t.Errorf("Read(%q) = %s, want %s", rec2.Name, string(contents2), contents)
	}
}

func TestReadWrite(t *testing.T) {
	r := Newc.Reader(bytes.NewReader(testCPIO))
	files, err := ReadAllRecords(r)
	if err != nil {
		t.Fatalf("Reading testCPIO reader: %v", err)
	}

	buf := &bytes.Buffer{}
	w := Newc.Writer(buf)
	if err := WriteRecords(w, files); err != nil {
		t.Fatalf("WriteRecords: %v", err)
	}

	if err := WriteTrailer(w); err != nil {
		t.Fatalf("WriteTrailer: %v", err)
	}

	r = Newc.Reader(bytes.NewReader(testCPIO))
	files, err = ReadAllRecords(r)
	if err != nil {
		t.Fatalf("Reading testCPIO reader: %v", err)
	}

	r = Newc.Reader(bytes.NewReader(buf.Bytes()))
	filesReadBack, err := ReadAllRecords(r)
	if err != nil {
		t.Fatalf("TestReadWrite: reading generated data: %v", err)
	}

	// Now check a few things: arrays should be same length, Headers should match,
	// names should be the same, and data should be the same. If this all works,
	// it means we read in serialized data, wrote it out, read it in, and the
	// structs all matched.
	if len(files) != len(filesReadBack) {
		t.Fatalf("[]file len from testCPIO %v and generated %v are not the same and should be", len(files), len(filesReadBack))
	}
	for i := range files {
		f1 := files[i]
		f2 := filesReadBack[i]

		if f1.Info != f2.Info {
			t.Errorf("index %d: testCPIO Info\n%v\ngenerated Info\n%v\n", i, f1.Info, f2.Info)
		}

		contents1, err := ioutil.ReadAll(uio.Reader(f1))
		if err != nil {
			t.Errorf("index %d(%q): can't read from the source: %v", i, f1.Name, err)
		}
		contents2, err := ioutil.ReadAll(uio.Reader(f2))
		if err != nil {
			t.Errorf("index %d(%q): can't read from the dest: %v", i, f2.Name, err)
		}
		if !bytes.Equal(contents1, contents2) {
			t.Errorf("index %d content: file 1 (%q) is %v, file 2 (%q) wanted %v", i, f1.Name, contents1, f2.Name, contents2)
		}
	}
}

func TestBad(t *testing.T) {
	r := Newc.Reader(bytes.NewReader(badCPIO))
	if _, err := r.ReadRecord(); err != io.EOF {
		t.Errorf("ReadRecord(badCPIO) got %v, want %v", err, io.EOF)
	}

	r = Newc.Reader(bytes.NewReader(badMagicCPIO))
	if _, err := r.ReadRecord(); err == nil {
		t.Errorf("Wanted bad magic err, got nil")
	}
}

/*
drwxrwxr-x   9 rminnich rminnich        0 Jan 22 22:18 .
drwxr-xr-x   2 root     root            0 Jan 22 22:18 etc
-rw-r--r--   1 root     root          118 Jan 22 22:18 etc/localtime
-rw-r--r--   1 root     root           81 Jan 22 22:18 etc/resolv.conf
drwxr-xr-x   2 root     root            0 Jan 22 22:18 lib64
drwxr-xr-x   2 root     root            0 Jan 22 22:18 tcz
drwxr-xr-x   2 root     root            0 Jan 22 22:18 bin
drwxr-xr-x   2 root     root            0 Jan 22 22:18 tmp
drwxr-xr-x   2 root     root            0 Jan 22 22:18 dev
crw-r--r--   1 root     root       5,   1 Jan 22 22:18 dev/console
crw-r--r--   1 root     root       4,  64 Jan 22 22:18 dev/ttyS0
brw-rw----   1 root     root       7,   2 Jan 22 22:18 dev/loop2
crw-------   1 root     root      10, 237 Jan 22 22:18 dev/loop-control
brw-rw----   1 root     root       7,   7 Jan 22 22:18 dev/loop7
brw-rw----   1 root     root       7,   6 Jan 22 22:18 dev/loop6
brw-rw----   1 root     root       7,   4 Jan 22 22:18 dev/loop4
brw-rw----   1 root     root       7,   1 Jan 22 22:18 dev/loop1
brw-rw----   1 root     root       7,   5 Jan 22 22:18 dev/loop5
crw-r--r--   1 root     root       1,   3 Jan 22 22:18 dev/null
brw-rw----   1 root     root       7,   0 Jan 22 22:18 dev/loop0
brw-rw----   1 root     root       7,   3 Jan 22 22:18 dev/loop3
drwxr-xr-x   3 root     root            0 Jan 22 22:18 usr
drwxr-xr-x   2 root     root            0 Jan 22 22:18 usr/lib
*/
var (
	badCPIO      = []byte{}
	badMagicCPIO = []byte{0, 0, 0, 0, 0, 0}
	testCPIO     = []byte{
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33,
		0x43, 0x31, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31, 0x46, 0x44, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x33, 0x45, 0x38, 0x30, 0x30, 0x30, 0x30, 0x30, 0x33,
		0x45, 0x38, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x39, 0x35, 0x38,
		0x38, 0x35, 0x41, 0x30, 0x34, 0x45, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x2e, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x45, 0x33, 0x43, 0x35, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31,
		0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x32, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x65, 0x74, 0x63, 0x00, 0x00, 0x00,
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x43, 0x32,
		0x30, 0x43, 0x30, 0x30, 0x30, 0x30, 0x38, 0x31, 0x41, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x35, 0x38,
		0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x37, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x45, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x65, 0x74, 0x63, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x74,
		0x69, 0x6d, 0x65, 0x00, 0x54, 0x5a, 0x69, 0x66, 0x32, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x54, 0x43, 0x00, 0x00, 0x00,
		0x54, 0x5a, 0x69, 0x66, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x55, 0x54, 0x43, 0x00, 0x00, 0x00, 0x0a, 0x55, 0x54, 0x43,
		0x30, 0x0a, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x37, 0x42, 0x30, 0x44, 0x30, 0x30, 0x30, 0x30, 0x38, 0x31,
		0x41, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x35, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x65, 0x74, 0x63, 0x2f, 0x72, 0x65,
		0x73, 0x6f, 0x6c, 0x76, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x00,
		0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x31,
		0x39, 0x32, 0x2e, 0x31, 0x36, 0x38, 0x2e, 0x31, 0x2e, 0x31, 0x30, 0x0a,
		0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x2e,
		0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x73, 0x69, 0x6e,
		0x67, 0x6c, 0x65, 0x2d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
		0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x3a, 0x31, 0x20, 0x61, 0x74,
		0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x3a, 0x35, 0x0a, 0x00, 0x00, 0x00,
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33,
		0x43, 0x42, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31, 0x45, 0x44, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x35, 0x38,
		0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x6c, 0x69, 0x62, 0x36, 0x34, 0x00, 0x30, 0x37, 0x30, 0x37,
		0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33, 0x43, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x34, 0x31, 0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30,
		0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x74, 0x63,
		0x7a, 0x00, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x45, 0x33, 0x43, 0x43, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31,
		0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x32, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x62, 0x69, 0x6e, 0x00, 0x00, 0x00,
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33,
		0x43, 0x44, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31, 0x45, 0x44, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x35, 0x38,
		0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x74, 0x6d, 0x70, 0x00, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37,
		0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33, 0x43, 0x36, 0x30, 0x30,
		0x30, 0x30, 0x34, 0x31, 0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30,
		0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65,
		0x76, 0x00, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x32, 0x31, 0x30, 0x30, 0x30, 0x30, 0x32, 0x31,
		0x41, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6f,
		0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x00, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37,
		0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x43, 0x43, 0x32, 0x32, 0x30, 0x30,
		0x30, 0x30, 0x32, 0x31, 0x41, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30,
		0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65,
		0x76, 0x2f, 0x74, 0x74, 0x79, 0x53, 0x30, 0x00, 0x30, 0x37, 0x30, 0x37,
		0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x43, 0x43, 0x32, 0x33, 0x30, 0x30,
		0x30, 0x30, 0x36, 0x31, 0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30,
		0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65,
		0x76, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x32, 0x00, 0x30, 0x37, 0x30, 0x37,
		0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x43, 0x43, 0x31, 0x46, 0x30, 0x30,
		0x30, 0x30, 0x32, 0x31, 0x38, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30,
		0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65,
		0x76, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
		0x6f, 0x6c, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x32, 0x30, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x37, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x37, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x36, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x42, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x34, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x44, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x31, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x45, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x35, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x32, 0x34, 0x30, 0x30, 0x30, 0x30, 0x32, 0x31,
		0x41, 0x34, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x33, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x39, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x75,
		0x6c, 0x6c, 0x00, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x38, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x30, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x43, 0x43, 0x31, 0x43, 0x30, 0x30, 0x30, 0x30, 0x36, 0x31,
		0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x31, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x37, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x33, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x6f,
		0x6f, 0x70, 0x33, 0x00, 0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30,
		0x33, 0x30, 0x45, 0x33, 0x43, 0x39, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31,
		0x45, 0x44, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x33, 0x35, 0x38, 0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x46, 0x43, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x34, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x75, 0x73, 0x72, 0x00, 0x00, 0x00,
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x33, 0x30, 0x45, 0x33,
		0x43, 0x41, 0x30, 0x30, 0x30, 0x30, 0x34, 0x31, 0x45, 0x44, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x35, 0x38,
		0x38, 0x35, 0x41, 0x30, 0x34, 0x41, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x46, 0x43, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x38, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x75, 0x73, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x00, 0x00, 0x00,
		0x30, 0x37, 0x30, 0x37, 0x30, 0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x42, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x54, 0x52, 0x41, 0x49, 0x4c, 0x45, 0x52, 0x21, 0x21, 0x21,
		0, 0, 0, 0,
	}

	testResult = []string{
		".: Ino 3204033 Mode 040775 UID 1000 GID 1000 NLink 9 MTime 2017-01-23 06:18:54 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"etc: Ino 3204037 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"etc/localtime: Ino 3195404 Mode 0100644 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 118 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"etc/resolv.conf: Ino 3177229 Mode 0100644 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 81 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"lib64: Ino 3204043 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"tcz: Ino 3204036 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"bin: Ino 3204044 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"tmp: Ino 3204045 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"dev: Ino 3204038 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"dev/console: Ino 3197985 Mode 020644 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 5 Rminor 1",
		"dev/ttyS0: Ino 3197986 Mode 020644 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 4 Rminor 64",
		"dev/loop2: Ino 3197987 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 2",
		"dev/loop-control: Ino 3197983 Mode 020600 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 10 Rminor 237",
		"dev/loop7: Ino 3197984 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 7",
		"dev/loop6: Ino 3197975 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 6",
		"dev/loop4: Ino 3197979 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 4",
		"dev/loop1: Ino 3197981 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 1",
		"dev/loop5: Ino 3197982 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 5",
		"dev/null: Ino 3197988 Mode 020644 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 1 Rminor 3",
		"dev/loop0: Ino 3197976 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 0",
		"dev/loop3: Ino 3197980 Mode 060660 UID 0 GID 0 NLink 1 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 7 Rminor 3",
		"usr: Ino 3204041 Mode 040755 UID 0 GID 0 NLink 3 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
		"usr/lib: Ino 3204042 Mode 040755 UID 0 GID 0 NLink 2 MTime 2017-01-23 06:18:50 +0000 UTC FileSize 0 Major 252 Minor 0 Rmajor 0 Rminor 0",
	}
)

// testReproducible verifies that we can produce reproducible cpio archives for newc format.
func TestReproducible(t *testing.T) {
	contents := []byte("LANAAAAAAAAAA")
	rec := []Record{StaticRecord(contents, Info{
		Ino:      1,
		Mode:     syscall.S_IFREG | 2,
		UID:      3,
		GID:      4,
		NLink:    5,
		MTime:    6,
		FileSize: 7,
		Major:    8,
		Minor:    9,
		Rmajor:   10,
		Rminor:   11,
		Name:     "foobar",
	}),
	}

	// First test that it fails unless we make it reproducible

	b1 := &bytes.Buffer{}
	w := Newc.Writer(b1)
	if err := WriteRecords(w, rec); err != nil {
		t.Errorf("Could not write record %q: %v", rec[0].Name, err)
	}
	rec[0].ReaderAt = bytes.NewReader(contents)
	b2 := &bytes.Buffer{}
	w = Newc.Writer(b2)
	rec[0].MTime++
	if err := WriteRecords(w, rec); err != nil {
		t.Errorf("Could not write record %q: %v", rec[0].Name, err)
	}

	if reflect.DeepEqual(b1.Bytes()[:], b2.Bytes()[:]) {
		t.Error("Reproducible: compared as same, wanted different")
	}

	// Second test that it works if we make it reproducible
	// It does indeed fail without the second call.

	b1 = &bytes.Buffer{}
	w = Newc.Writer(b1)
	rec[0].ReaderAt = bytes.NewReader([]byte(contents))
	MakeAllReproducible(rec)
	if err := WriteRecords(w, rec); err != nil {
		t.Errorf("Could not write record %q: %v", rec[0].Name, err)
	}

	b2 = &bytes.Buffer{}
	w = Newc.Writer(b2)
	rec[0].MTime++
	rec[0].ReaderAt = bytes.NewReader([]byte(contents))
	MakeAllReproducible(rec)
	if err := WriteRecords(w, rec); err != nil {
		t.Errorf("Could not write record %q: %v", rec[0].Name, err)
	}

	if len(b1.Bytes()) != len(b2.Bytes()) {
		t.Fatalf("Reproducible \n%v,\n%v: len is different, wanted same", b1.Bytes()[:], b2.Bytes()[:])
	}
	if !reflect.DeepEqual(b1.Bytes()[:], b2.Bytes()[:]) {
		t.Error("Reproducible: compared different, wanted same")
		for i := range b1.Bytes() {
			a := b1.Bytes()[i]
			b := b2.Bytes()[i]
			if a != b {
				t.Errorf("\tb1[%d] is %v, b2[%d] is %v", i, a, i, b)
			}
		}
	}
}
