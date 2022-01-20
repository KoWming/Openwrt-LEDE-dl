// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// +build debug

package drpcdebug

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var logger = log.New(os.Stderr, "", 0)

// Log executes the callback for a string to log if built with the debug tag.
func Log(cb func() (who, what, why string)) {
	_, file, line, _ := runtime.Caller(1)
	where := fmt.Sprintf("%s:%d", filepath.Base(file), line)
	who, what, why := cb()
	logger.Output(2, fmt.Sprintf("%24s | %-26s | %-6s | %s",
		where, who, what, why))
}
