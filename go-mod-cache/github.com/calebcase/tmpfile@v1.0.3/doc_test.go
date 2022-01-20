// Copyright 2019 Caleb Case

package tmpfile

func ExampleNew_default() {
	// Use the system default temp directory (as returned by os.TempDir().
	// This is equivalent to New(os.TempDir(), "example-").
	f, err := New("", "example-")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func ExampleNew_dir() {
	// Use a local directory for the temporary files. This directory must
	// already exist.
	f, err := New("ephemeral", "example-")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
