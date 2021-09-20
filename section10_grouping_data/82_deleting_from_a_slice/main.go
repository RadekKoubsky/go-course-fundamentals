package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	// func append(slice []Type, elems ...Type) []Type
	x = append(x, 77, 88, 99)
	fmt.Println(x)

	y := []int{234, 456, 789}
	x = append(x, y...)
	fmt.Println(x)

	// deleting from slice using ":" operator, see section "80_slicing_a_slice"
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
