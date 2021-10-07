package main

import "fmt"

/**
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors

Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

*/

type person struct {
	first string
	last string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		favoriteFlavors: []string{"Lemon", "Chocolate"},
	}

	fmt.Println(p1.first, p1.last)
	for _, flavor := range p1.favoriteFlavors {
		fmt.Println(flavor)
	}
}
