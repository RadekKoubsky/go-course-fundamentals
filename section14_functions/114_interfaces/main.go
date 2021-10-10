package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I wass passed into barrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I wass passed into barrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed in to bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Doctor",
		last:  "No",
	}
	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
