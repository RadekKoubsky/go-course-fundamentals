package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo()
	fmt.Printf("Sum of %v = %v", x, sum)
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

