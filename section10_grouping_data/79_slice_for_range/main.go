package main

import "fmt"

func main() {
	x := []int{4,5,6,7,42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])

	/*
	Iteration values for array and slice in for range loop are:
	Range expression							1st value			2nd value
	array or slice  a  [n]E, *[n]E, or []E		index 	i int		a[i]	E
	*/
	for index, value := range x{
		fmt.Printf("index %d, value %d\n", index, value)
	}
}
