// Copyright 2020 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package core_test

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/stretchr/testify/assert"
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
)

type mockNotifier struct {
	actions []string
	lock    sync.Mutex
}

func (m *mockNotifier) BeforeLoginUser(conn *core.Conn, userName string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeLoginUser")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforePutFile(conn *core.Conn, dstPath string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforePutFile")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforeDeleteFile(conn *core.Conn, dstPath string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeDeleteFile")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforeChangeCurDir(conn *core.Conn, oldCurDir, newCurDir string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeChangeCurDir")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforeCreateDir(conn *core.Conn, dstPath string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeCreateDir")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforeDeleteDir(conn *core.Conn, dstPath string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeDeleteDir")
	m.lock.Unlock()
}
func (m *mockNotifier) BeforeDownloadFile(conn *core.Conn, dstPath string) {
	m.lock.Lock()
	m.actions = append(m.actions, "BeforeDownloadFile")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterUserLogin(conn *core.Conn, userName, password string, passMatched bool, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterUserLogin")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterFilePut(conn *core.Conn, dstPath string, size int64, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterFilePut")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterFileDeleted(conn *core.Conn, dstPath string, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterFileDeleted")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterCurDirChanged(conn *core.Conn, oldCurDir, newCurDir string, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterCurDirChanged")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterDirCreated(conn *core.Conn, dstPath string, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterDirCreated")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterDirDeleted(conn *core.Conn, dstPath string, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterDirDeleted")
	m.lock.Unlock()
}
func (m *mockNotifier) AfterFileDownloaded(conn *core.Conn, dstPath string, size int64, err error) {
	m.lock.Lock()
	m.actions = append(m.actions, "AfterFileDownloaded")
	m.lock.Unlock()
}

func assetMockNotifier(t *testing.T, mock *mockNotifier, lastActions []string) {
	if len(lastActions) == 0 {
		return
	}
	mock.lock.Lock()
	assert.EqualValues(t, lastActions, mock.actions[len(mock.actions)-len(lastActions):])
	mock.lock.Unlock()
}

func TestNotification(t *testing.T) {
	err := os.MkdirAll("./testdata", os.ModePerm)
	assert.NoError(t, err)

	var perm = core.NewSimplePerm("test", "test")
	opt := &core.ServerOpts{
		Name: "test ftpd",
		Factory: &file.DriverFactory{
			RootPath: "./testdata",
			Perm:     perm,
		},
		Port: 2121,
		Auth: &core.SimpleAuth{
			Name:     "admin",
			Password: "admin",
		},
		Logger: new(core.DiscardLogger),
	}

	mock := &mockNotifier{}

	runServer(t, opt, []core.Notifier{mock}, func() {
		// Give server 0.5 seconds to get to the listening state
		timeout := time.NewTimer(time.Millisecond * 500)

		for {
			f, err := ftp.Connect("localhost:2121")
			if err != nil && len(timeout.C) == 0 { // Retry errors
				continue
			}
			assert.NoError(t, err)

			assert.NoError(t, f.Login("admin", "admin"))
			assetMockNotifier(t, mock, []string{"BeforeLoginUser", "AfterUserLogin"})

			assert.Error(t, f.Login("admin", "1111"))
			assetMockNotifier(t, mock, []string{"BeforeLoginUser", "AfterUserLogin"})

			var content = `test`
			assert.NoError(t, f.Stor("server_test.go", strings.NewReader(content)))
			assetMockNotifier(t, mock, []string{"BeforePutFile", "AfterFilePut"})

			r, err := f.RetrFrom("/server_test.go", 2)
			assert.NoError(t, err)

			buf, err := ioutil.ReadAll(r)
			r.Close()
			assert.NoError(t, err)
			assert.EqualValues(t, "st", string(buf))
			assetMockNotifier(t, mock, []string{"BeforeDownloadFile", "AfterFileDownloaded"})

			err = f.Rename("/server_test.go", "/test.go")
			assert.NoError(t, err)

			err = f.MakeDir("/src")
			assert.NoError(t, err)
			assetMockNotifier(t, mock, []string{"BeforeCreateDir", "AfterDirCreated"})

			err = f.Delete("/test.go")
			assert.NoError(t, err)
			assetMockNotifier(t, mock, []string{"BeforeDeleteFile", "AfterFileDeleted"})

			err = f.ChangeDir("/src")
			assert.NoError(t, err)
			assetMockNotifier(t, mock, []string{"BeforeChangeCurDir", "AfterCurDirChanged"})

			err = f.RemoveDir("/src")
			assert.NoError(t, err)
			assetMockNotifier(t, mock, []string{"BeforeDeleteDir", "AfterDirDeleted"})

			err = f.Quit()
			assert.NoError(t, err)

			break
		}
	})
}
