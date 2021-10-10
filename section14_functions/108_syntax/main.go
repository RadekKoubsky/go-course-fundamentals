package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

// func (r receiver) identifier(parameters) (return(s)) { ... }
/*
 Different wording: "parameters vs arguments"

 We define our func with parameters (if any)
 We call our func and pass in arguments (if any)
*/

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello", s)
}

// returns 1 value
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

// returns 2 values
func mouse(firstName string, lastName string) (string, bool) {
	a := fmt.Sprint(firstName, " ", lastName, `, says "Hello"`)
	b := false
	 return a, b
}
