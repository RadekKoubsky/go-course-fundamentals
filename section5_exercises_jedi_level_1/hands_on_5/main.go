package main

import "fmt"

type anotherInt int

var y int
func main() {
	var x anotherInt
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
