package main

import "fmt"

type anotherInt int

func main() {
	var x anotherInt
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
