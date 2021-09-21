package main

import "fmt"

/*
A slice created with make always allocates a new, hidden array to which the returned slice value refers.

If you know the size of a slice upfront, use "make" built-in function to init the slice with an underlying array of that size,
thus the array does not have to be recomputed every time it gets full.
*/
func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[10] = 3423 throws index out of range, need to use append
	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// now the array is full, its capacity will be doubled next time we add element to slice
	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// now the capacity of the array is doubled to accommodate the next element
	fmt.Println("Doubling capacity of array from", cap(x))
	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println("Length", len(x))
	fmt.Println("Capacity", cap(x))
}
