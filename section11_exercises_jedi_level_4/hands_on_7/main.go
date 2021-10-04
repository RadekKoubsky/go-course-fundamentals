package main

import "fmt"

/**
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
*/
func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James"}}

	for _, strings := range x {
		fmt.Println("Nested slice content:")
		for _, s := range strings {
			fmt.Println("\t", s)
		}
	}
}
