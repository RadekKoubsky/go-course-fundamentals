package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

// NON-POINTER RECEIVER
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// POINTER RECEIVER
func (s *square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	fmt.Println("area", s.area())
}

/*
a NON-POINTER RECEIVER
works with values that are POINTERS or NON-POINTERS.

a POINTER RECEIVER
only works with values that are POINTERS.

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

*/
func main() {
	// NON-POINTER RECEIVER - works with values that are POINTERS or NON-POINTERS
	c := circle{5}
	// NON-POINTER RECEIVER & NON-POINTER VALUE
	info(c)
	// NON-POINTER RECEIVER & POINTER VALUE
	info(&c)

	//POINTER RECEIVER - only works with values that are POINTERS
	s := square{2}
	// POINTER RECEIVER & POINTER VALUE
	info(&s)

	// POINTER RECEIVER & NON-POINTER VALUE
	// info(s) does not compile
}
