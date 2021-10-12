package main

import (
	"fmt"
)

/*
Callback: passing a func as an argument to call it later
*/
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Printf("Sum of %v: %v\n", ii, s)

	s2 := evenSum(sum, ii...)
	fmt.Printf("Sum of even numbers in %v: %v\n", ii, s2)

	s3 := oddSum(sum, ii...)
	fmt.Printf("Sum of odd numbers in %v: %v\n", ii, s3)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if (v %2 == 0){
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func oddSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if (v %2 != 0){
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
