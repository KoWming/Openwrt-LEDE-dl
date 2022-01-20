// Create a temporary file in the local directory.

package main

import "github.com/calebcase/tmpfile"

func main() {
	f, err := tmpfile.New("", "example-*")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Example Data")
}
