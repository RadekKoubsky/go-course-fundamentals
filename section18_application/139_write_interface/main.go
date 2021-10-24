package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello world")
	// os.Stdout is of type File,
	// the File type has Write(p []byte) method attached to it,
	// Write interface defines Write(p []byte) method
	// this means File is also of type "Writer" and
	// we can pass os.Stdout to Fprintln(w io.Writer, a ...interface{}) method
	// which accepts Writer interface
	fmt.Fprintln(os.Stdout, "Hello world")
	io.WriteString(os.Stdout, "Hello world")
}
