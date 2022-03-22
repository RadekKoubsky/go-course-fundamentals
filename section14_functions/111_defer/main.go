package main

import "fmt"

func main() {
	/*
		A "defer" statement invokes a function whose execution is deferred to the
		moment the surrounding function returns, either because the surrounding
		function executed a return statement, reached the end of its function body,
		or because the corresponding goroutine is panicking.

		Typical use case:
		Opening a file, do something, close file at the end.
		With defer, the closeFile function will be called whenever
		the surrounding function exits, thus we do not have to care
		about where to place the closing code.

		More on defer at https://go.dev/blog/defer-panic-and-recover
	*/
	// foo will be executed when the main method exits, thus after bar() call
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
