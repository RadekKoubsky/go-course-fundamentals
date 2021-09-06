package main

import "fmt"

const (
	a = 1
	b = 2.34
	c = "James Bond"
)

const (
	x int = 42
	y string = "Hello"
	z float64 = 42.42
)

func main() {
	fmt.Printf("%v\t\t %T\n", a, a)
	fmt.Printf("%v\t\t %T\n", b, b)
	fmt.Printf("%v\t %T\n", c, c)
	fmt.Printf("%v\t\t %T\n", x, x)
	fmt.Printf("%v\t\t %T\n", y, y)
	fmt.Printf("%v\t\t %T\n", z, z)
}