package main

import "fmt"

/*
Two types of constants:
typed, untyped

Untyped constants give compiler more flexibility, it can assign it to another types
*/
const (
	A int 		= 42
	B float64 	= 42.78
	C 			= "James Bond"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Printf("%T\n", A)
	fmt.Printf("%T\n", B)
	fmt.Printf("%T\n", C)
}
