package main

import "fmt"

// the type int is inferred
var x = 34

// name is of type string
// Go is statically-typed language
var name string = "John"

var rawStringText string = `This is raw string as it is, "I can use double quotes and

new line and anything else"`

func main() {
	fmt.Println(x)
	fmt.Printf("Type of x is %T\n", x)

	fmt.Println(name)
	fmt.Printf("Type of name is %T\n", name)

	fmt.Println(rawStringText)

	myFormattedString, _ := fmt.Printf("My formatted string %d", 45)
	fmt.Println(myFormattedString)
}


