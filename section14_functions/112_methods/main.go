package main

import (
	"fmt"
)

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

// func speak() can be attached to any value of type secretAgent as defined in the receiver clause
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// func (r receiver) identifier(parameters) (returns(s)) { ... }

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}
