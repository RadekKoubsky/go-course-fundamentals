package main

import "fmt"

func main() {
	fallThrough()
	defaultExample()
	switchOnValue()
	switchOnMultipleValues()
}

func switchOnMultipleValues() {
	fmt.Println("Switch on multiple values in case")
	name := "Bond"
	switch name {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("match money or bond or dr no")
	case "M":
		fmt.Println("M")
	default:
		fmt.Println("this is default")
	}
}

func switchOnValue() {
	fmt.Println("Switch on value example")
	name := "Bond"
	switch name {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("james bond")
	default:
		fmt.Println("this is default")
	}
}

func defaultExample() {
	fmt.Println("Default example")
	// missing switch expression defaults to "true"
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	default:
		fmt.Println("this is default")
	}
}

func fallThrough() {
	fmt.Println("Fallthrough example")
	// missing switch expression defaults to "true"
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		// there is no default fallthrough, we must explicitly write it
		fmt.Println("this should print")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
		// fallthrough always execute the next case statement either its true or false
		// thus the nexe case will be executed even though is false and so on
		// using fallthrough this is not recommended
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
	default:
		// executes only if nothing was true
		// or use fallthrough in previous case to always execute default
		fmt.Println("this is default")
	}
}
