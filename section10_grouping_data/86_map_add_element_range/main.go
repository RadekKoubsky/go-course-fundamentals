package main

import "fmt"

func main() {
	capitalsByCountry := map[string]string{
		"US": "Washington D.C.",
		"UK": "Longon",
	}
	fmt.Println(capitalsByCountry)

	// adds element to map
	capitalsByCountry["Netherlands"] = "Amsterdam"
	fmt.Println(capitalsByCountry)

	/*
		Iteration values for map in for range loop are:
		Range expression							1st value			2nd value
		map  m  map[K]V                				key  k  K       	m[k]   V
	*/
	for key, value := range capitalsByCountry {
		fmt.Printf("%s=%s\n", key, value)
	}
}
