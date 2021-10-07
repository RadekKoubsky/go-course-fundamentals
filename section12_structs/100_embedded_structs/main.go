package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

/*
A field declared with a type but no explicit field name is called an embedded field.
An embedded field must be specified as a type name T or as a pointer to a non-interface
type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.
*/

type secretAgent struct {
	person // an embedded struct, embedded field
	first string
	ltk bool
}

func main() {
	secretAgent := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age: 32,
		},
		first: "something collision",
		ltk: true,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age: 27,
	}

	fmt.Println(secretAgent)
	fmt.Println(p2)

	// secretAgent.last -> the inner type gets promoted to the outer type, thus we can access inner.last with outer.last
	fmt.Println(secretAgent.first, secretAgent.person.first, secretAgent.last, secretAgent.age, secretAgent.ltk)

	fmt.Println(p2.first, p2.last, p2.age)
}
