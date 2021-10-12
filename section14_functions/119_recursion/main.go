package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorialLoop(4))
}

func factorial(n int) int {
	if (n == 0) {
		return 1;
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	fact := 1
	for ; n > 1; n-- {
		fact = fact * n
	}
	return fact
}
