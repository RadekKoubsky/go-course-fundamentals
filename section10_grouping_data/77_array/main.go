package main

import "fmt"

/*
Effective go on arrays: https://golang.org/doc/effective_go#arrays

Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation,
but primarily they are a building block for slices, the subject of the next section.

Use slices instead.
*/
func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
