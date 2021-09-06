package main

import "fmt"

func main() {
	x := 1
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%#X\n", y, y, y)

	z := x << 2
	fmt.Printf("%d\t%b\t%#X", z, z, z)
}
