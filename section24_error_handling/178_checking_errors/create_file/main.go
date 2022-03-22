package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)
}
