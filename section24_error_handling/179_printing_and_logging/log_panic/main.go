package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		/*
			When a function F calls panic, normal execution of F stops immediately.
			Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.
		*/
		defer foo()

		/*
			Panicln is equivalent to Println() followed by a call to panic().

			Panics are similar to C++ and Java exceptions, but are only intended for run-time errors,
			such as following a nil pointer or attempting to index an array out of bounds.

			A panic stops the normal execution of a goroutine:

			When a program panics, it immediately starts to unwind the call stack.
			This continues until the program crashes and prints a stack trace,
			or until the built-in recover function is called.

			A panic is caused either by a runtime error, or an explicit call to the built-in panic function.

			Blog post about panic https://go.dev/blog/defer-panic-and-recover
		*/
		log.Panicln(err)
	}
}

func foo() {
	fmt.Println("Panic allows deferred functions to run")
}
