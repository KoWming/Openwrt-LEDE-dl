// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package core_test

import (
	"io/ioutil"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/stretchr/testify/assert"
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
)

func runServer(t *testing.T, opt *core.ServerOpts, notifiers []core.Notifier, execute func()) {
	s := core.NewServer(opt)
	for _, notifier := range notifiers {
		s.RegisterNotifer(notifier)
	}
	go func() {
		err := s.ListenAndServe()
		assert.EqualError(t, err, core.ErrServerClosed.Error())
	}()

	execute()

	assert.NoError(t, s.Shutdown())
}

func TestFileDriver(t *testing.T) {
	err := os.MkdirAll("./testdata", os.ModePerm)
	assert.NoError(t, err)

	var perm = core.NewSimplePerm("test", "test")
	opt := &core.ServerOpts{
		Name: "test ftpd",
		Factory: &file.DriverFactory{
			RootPath: "./testdata",
			Perm:     perm,
		},
		Port: 2122,
		Auth: &core.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Logger: new(core.DiscardLogger),
	}

	runServer(t, opt, nil, func() {
		// Give server 0.5 seconds to get to the listening state
		timeout := time.NewTimer(time.Millisecond * 500)

		for {
			f, err := ftp.Connect("localhost:2122")
			if err != nil && len(timeout.C) == 0 { // Retry errors
				continue
			}
			assert.NoError(t, err)

			assert.NoError(t, f.Login("admin", "admin"))
			assert.Error(t, f.Login("admin", ""))

			var content = `test`
			assert.NoError(t, f.Stor("server_test.go", strings.NewReader(content)))

			names, err := f.NameList("/")
			assert.NoError(t, err)
			assert.EqualValues(t, 1, len(names))
			assert.EqualValues(t, "server_test.go", names[0])

			bs, err := ioutil.ReadFile("./testdata/server_test.go")
			assert.NoError(t, err)
			assert.EqualValues(t, content, string(bs))

			entries, err := f.List("/")
			assert.NoError(t, err)
			assert.EqualValues(t, 1, len(entries))
			assert.EqualValues(t, "server_test.go", entries[0].Name)
			assert.EqualValues(t, 4, entries[0].Size)
			assert.EqualValues(t, ftp.EntryTypeFile, entries[0].Type)

			curDir, err := f.CurrentDir()
			assert.NoError(t, err)
			assert.EqualValues(t, "/", curDir)

			size, err := f.FileSize("/server_test.go")
			assert.NoError(t, err)
			assert.EqualValues(t, 4, size)

			r, err := f.RetrFrom("/server_test.go", 2)
			assert.NoError(t, err)

			buf, err := ioutil.ReadAll(r)
			r.Close()
			assert.NoError(t, err)
			assert.EqualValues(t, "st", string(buf))

			err = f.Rename("/server_test.go", "/test.go")
			assert.NoError(t, err)

			err = f.MakeDir("/src")
			assert.NoError(t, err)

			err = f.Delete("/test.go")
			assert.NoError(t, err)

			err = f.ChangeDir("/src")
			assert.NoError(t, err)

			curDir, err = f.CurrentDir()
			assert.NoError(t, err)
			assert.EqualValues(t, "/src", curDir)

			assert.NoError(t, f.Stor("server_test.go", strings.NewReader(content)))

			r, err = f.Retr("/src/server_test.go")
			assert.NoError(t, err)

			buf, err = ioutil.ReadAll(r)
			r.Close()
			assert.NoError(t, err)
			assert.EqualValues(t, "test", string(buf))

			err = f.RemoveDir("/src")
			assert.NoError(t, err)

			err = f.Quit()
			assert.NoError(t, err)

			break
		}
	})
}

func TestLogin(t *testing.T) {
	err := os.MkdirAll("./testdata", os.ModePerm)
	assert.NoError(t, err)

	var perm = core.NewSimplePerm("test", "test")

	// Server options without hostname or port
	opt := &core.ServerOpts{
		Name: "test ftpd",
		Factory: &file.DriverFactory{
			RootPath: "./testdata",
			Perm:     perm,
		},
		Auth: &core.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Logger: new(core.DiscardLogger),
	}

	// Start the listener
	l, err := net.Listen("tcp", ":2123")
	assert.NoError(t, err)

	// Start the server using the listener
	s := core.NewServer(opt)
	go func() {
		err := s.Serve(l)
		assert.EqualError(t, err, core.ErrServerClosed.Error())
	}()

	// Give server 0.5 seconds to get to the listening state
	timeout := time.NewTimer(time.Millisecond * 500)
	for {
		f, err := ftp.Connect("localhost:2123")
		if err != nil && len(timeout.C) == 0 { // Retry errors
			continue
		}
		assert.NoError(t, err)

		assert.NoError(t, f.Login("admin", "admin"))
		assert.Error(t, f.Login("admin", ""))

		err = f.Quit()
		assert.NoError(t, err)
		break
	}

	assert.NoError(t, s.Shutdown())
}
