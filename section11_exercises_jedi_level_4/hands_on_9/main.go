package main

import "fmt"

/**
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop


*/
func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["homer_simpson"] = []string{`Donuts`, `TV`, `Beer`}

	for person, favoriteThings := range m {
		fmt.Printf("%v favorite things:\n", person)
		for i, favoriteThing := range favoriteThings {
			fmt.Printf("\t%v %v\n", i, favoriteThing)
		}
		fmt.Println()
	}
}
