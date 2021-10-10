package main

import "fmt"

func main() {
	// bar() returns a func, func is assigned to x
	x := bar()
	fmt.Printf("func bar() returned '%T'\n", x)

	// execute x which is func returned from bar() call above
	fmt.Println("Called func returned from bar(), result: ", x())

	// we can call func returned by bar() directly after bar() call
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}
