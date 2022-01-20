// Copyright 2019 Caleb Case

package tmpfile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	var err error

	for _, tc := range []struct {
		Name      string
		TmpDir    string
		TmpPrefix string
		TmpSuffix string
	}{
		{"default dir with prefix", "", "tmp-test-", ""},
		{"default dir with prefix and suffix", "", "tmp-test-", "-tmp-test"},
		{"default dir with suffix", "", "", "-tmp-test"},
		{"cwd dir with prefix", ".", "tmp-test-", ""},
		{"cwd dir with prefix and suffix", ".", "tmp-test-", "-tmp-test"},
		{"cwd dir with suffix", ".", "", "-tmp-test"},
	} {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			var f *os.File

			name := strings.Join([]string{tc.TmpPrefix, tc.TmpSuffix}, "*")

			f, err = New(tc.TmpDir, name)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				err = f.Close()
				if err != nil {
					t.Fatal(err)
				}
			}()

			t.Logf("Created temporary file: %q\n", f.Name())

			// Temporary file name should match the pattern.
			basename := filepath.Base(f.Name())

			if tc.TmpPrefix != "" && !strings.HasPrefix(basename, tc.TmpPrefix) {
				t.Fatalf("Temporary file name does not match prefix: %q != %q\n", name, basename)
			}

			if tc.TmpSuffix != "" && !strings.HasSuffix(basename, tc.TmpSuffix) {
				t.Fatalf("Temporary file name does not match suffix: %q != %q\n", name, basename)
			}

			// Temporary file should not exist.
			err = notExists(f.Name())
			if err != nil {
				t.Fatalf("Temporary file exists, but it should have been unlinked already: %+v\n", err)
			}

			// Temporary file should still be read/write/seek-able.
			msg := "Hello world!\n"
			_, err = f.Write([]byte(msg))
			if err != nil {
				t.Fatal(err)
			}

			err = f.Sync()
			if err != nil {
				t.Fatal(err)
			}

			var offset int64

			offset, err = f.Seek(0, io.SeekStart)
			if err != nil {
				t.Fatal(err)
			}

			if offset != 0 {
				t.Fatalf("Seeking to 0 offset failed (still at %d).\n", offset)
			}

			var data []byte

			data, err = ioutil.ReadAll(f)
			if err != nil {
				t.Fatal(err)
			}

			if string(data) != msg {
				t.Fatalf("Reading data failed. %q != %q\n", msg, string(data))
			}
		})
	}
}

// This is a utility test whose purpose is to open a temporary file and then
// exit without cleaning it up.
func TestNoCleanup(t *testing.T) {
	var err error

	if _, ok := os.LookupEnv("TMP_TEST_EXEC"); !ok {
		t.Skip("skipping (only useful for exec testing")
	}

	var f *os.File

	f, err = New("", "test-no-cleanup-")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", f.Name())
}

// This is a utility test whose purpose is to open a temporary file and then
// hang without cleaning it up.
func TestPause(t *testing.T) {
	var err error

	if _, ok := os.LookupEnv("TMP_TEST_EXEC"); !ok {
		t.Skip("skipping (only useful for exec testing")
	}

	var f *os.File

	f, err = New("", "test-pause-")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Fprintf(os.Stderr, "%s\n", f.Name())

	c := make(chan bool)
	<-c
}

func TestExecKill(t *testing.T) {
	if _, ok := os.LookupEnv("TMP_TEST_EXEC"); ok {
		t.Skip("skipping (already in an exec test)")
	}

	var err error

	t.Log("Preparing command...")
	cmd := exec.Command("go", "test", "-run", "TestPause")
	cmd.Env = append(os.Environ(), "TMP_TEST_EXEC=true")

	t.Log("Setting up stdout pipe...")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	t.Log("Start command...", cmd)
	err = cmd.Start()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Read temp file path...")
	path, _, err := bufio.NewReader(stdout).ReadLine()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Temp file path: %q\n", string(path))

	t.Log("Kill command...")
	err = cmd.Process.Kill()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Wait for command...")
	err = cmd.Wait()
	if err == nil {
		t.Fatalf("Expected the process to exit with an error, but it didn't.\n")
	}

	t.Log("Check that the temp file is removed...")
	err = notExists(string(path))
	if err != nil {
		t.Fatalf("Temporary file exists, but it should have been unlinked already: %+v\n", err)
	}
}

func TestExecNoCleanup(t *testing.T) {
	if _, ok := os.LookupEnv("TMP_TEST_EXEC"); ok {
		t.Skip("skipping (already in an exec test)")
	}

	var err error

	t.Log("Preparing command...")
	cmd := exec.Command("go", "test", "-run", "TestNoCleanup")
	cmd.Env = append(os.Environ(), "TMP_TEST_EXEC=true")

	t.Log("Setting up stdout pipe...")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	t.Log("Start command...", cmd)
	err = cmd.Start()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Read temp file path...")
	path, _, err := bufio.NewReader(stdout).ReadLine()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Temp file path: %q\n", string(path))

	t.Log("Wait for command...")
	err = cmd.Wait()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Check that the temp file is removed...")
	err = notExists(string(path))
	if err != nil {
		t.Fatalf("Temporary file exists, but it should have been unlinked already: %+v\n", err)
	}
}
