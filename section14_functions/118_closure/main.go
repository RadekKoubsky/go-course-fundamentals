package main

import "fmt"

/*
Closure
Enclosing code around a variable, the scope is limited to the one area of code.
*/
func main() {
	// closure
	// This block is enclosing "y", limiting the scope of y
	{
	y := 42
	fmt.Println(y)
	}
	// does not compile, y is enclosed by block {}
	// fmt.Println(y)

	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // x =1
	fmt.Println(a()) // x= 2
	fmt.Println(a()) //x = 3

	// creating a new func incrementor and storing it to b encloses a new x variable, not the one used in a()
	// thus x starts from 0
	fmt.Println(b()) // x = 1
	fmt.Println(b()) // x = 2
	fmt.Println(b()) // x = 3
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
