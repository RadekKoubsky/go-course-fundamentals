package main

import "fmt"

var a int

// myInt type with underlying type int
type myInt int

var y myInt

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	y = 43
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// a = y cannot assign myInt type to int type

	// using conversion to assign a = y
	a = int(y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
