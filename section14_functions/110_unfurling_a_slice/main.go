package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(mySlice...)
	fmt.Printf("Sum of %v = %v", mySlice, sum)
	// variadic means 0 or more, thus we can pass 0 arguments to variadic func
	sumWithNoArgs := foo()
	fmt.Println("Sum with no args =", sumWithNoArgs)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}
