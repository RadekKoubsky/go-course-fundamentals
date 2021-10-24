package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// use json:jsonFieldName tags to map json property to struct field
	// by default, it uses implicit mapping struct field name <-> json field name
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

/*
Convert json to go struct https://mholt.github.io/json-to-go/
*/
func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	byteSlice := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", byteSlice)

	var people []person

	err := json.Unmarshal(byteSlice, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deserialized json to data:", people)
}
