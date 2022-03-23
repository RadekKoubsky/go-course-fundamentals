package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySum(1, 2, 3))
}

func mySum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
