package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	/*
	Slicing a slice using ":" colon operator
	x[1:3]
	Get values from index 1 (included) to index 3 (excluded)
	*/
	fmt.Println(x[1:3])

	for index, value := range x {
		fmt.Println(index, value)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
