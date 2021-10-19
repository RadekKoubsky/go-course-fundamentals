package main

import "fmt"

/**
When to use pointers?

1. You do not want to pass around big chunks of data e.g. data from database

2. You want to change value at some location/address, you want to mutate the value at the memory address, use pointers


*/
func main() {
	fmt.Println("Passing x param by value")
	x := 0
	fmt.Println("x befor", x)
	foo(x)
	fmt.Println("x after", x)

	fmt.Println("Passing x as pointer")
	x = 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	fooWithPoiner(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

// param int is passed by value
func foo(y int) {
	fmt.Println("y befor", y)
	y = 43
	fmt.Println("y after", y)
}

// param pointer *int is passed by value
func fooWithPoiner(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
