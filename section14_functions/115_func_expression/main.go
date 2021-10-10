package main

import "fmt"

func main() {
	// func is first class citizen in go, you can assign func to a variable etc.
	f := func() {
		fmt.Println("my first funct expression")
	}
	f()

	fWithParam := func(x int) {
		fmt.Println("My nubmer is: ", x)
	}
	fWithParam(1984)
}
