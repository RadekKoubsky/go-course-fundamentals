package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	// show address in memory where value of 'a' is stored
	fmt.Println(&a) // gives you the address

	fmt.Printf("%T\n", a)

	// prints pointer to an int '*int', pointer has its own type, 'int' is a different type than '*int'
	fmt.Printf("%T\n", &a)

	// Cannot use '&a' (type *int) as the type int
	// var b int = &a

	b := &a
	fmt.Println(b)

	// de-referencing an address, it returns the value stored at that address which is 42
	fmt.Println(*b)

	fmt.Println(*&a) // returns value stored at &a: 42
	fmt.Println(*&b) // returns value stored at &b: &a
	fmt.Println(**&b) // returns value stored at **&b: *(*&b) -> *(&a) -> 42

	*b = 43 // set value stored at address b (which is &a)
	fmt.Println(a) // now a = 43
}
