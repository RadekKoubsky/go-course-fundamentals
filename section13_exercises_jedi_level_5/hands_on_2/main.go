package main

import "fmt"

/**
Take the code from the previous exercise, then store the values of type
person in a map with the key of last name. Access each value in the map.
Print out the values, ranging over the slice

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

	p2 := person{
		first: "John",
		last: "Wayne",
		favoriteFlavors: []string{"Tequila", "Rum"},
	}

	myMap := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range myMap {
		fmt.Println(v.first, v.last)
		for _, flavor := range v.favoriteFlavors {
			fmt.Println(flavor)
		}
		fmt.Println("=======")
	}
}
