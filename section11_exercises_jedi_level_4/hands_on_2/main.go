package main

import "fmt"

/**
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES

Range over the slice and print the values out.

Using format printing
print out the TYPE of the slice

*/
func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slc {
		fmt.Printf("index %v value %v\n", i, v)
	}
	fmt.Printf("Type of slice: %T", slc)

}
