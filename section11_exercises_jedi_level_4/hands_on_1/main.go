package main

import "fmt"

/**
Hands-on exercise #1
Using a COMPOSITE LITERAL:

create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.

Range over the array and print the values out.

Using format printing
print out the TYPE of the array
*/
func main() {
	var arr = [5]int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Printf("index %v value %v\n", i, v)
	}
	fmt.Printf("Type of array: %T", arr)
}
