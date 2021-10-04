package main

import "fmt"

/**
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

append to that slice this value
52

print out the slice

in ONE STATEMENT append to that slice these values
53
54
55

print out the slice

append to the slice this slice
y := []int{56, 57, 58, 59, 60}
print out the slice

*/
func main() {
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	slc = append(slc, 52)
	fmt.Println(slc)

	slc = append(slc, 53, 54, 55)
	fmt.Println(slc)

	y := []int{56, 57, 58, 59, 60}
	slc = append(slc, y...)
	fmt.Println(slc)

}
