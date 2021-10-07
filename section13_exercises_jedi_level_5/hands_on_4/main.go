package main

import (
	"fmt"
)

/**
Create and use an anonymous struct.

*/
func main() {
	anonymousStruct := struct {
		countryByCity map[string]string
		languages     []string
	}{
		countryByCity: map[string]string{
			"Prague": "Czechia",
			"Los Angeles": "United states",
		},
		languages: []string{"CZ", "EN"},
	}

	for k, v := range anonymousStruct.countryByCity {
		fmt.Println(k, v)
	}

	for _, language := range anonymousStruct.languages {
		fmt.Println(language)
	}
}
