package main

import "fmt"

var y= 42

// fmt package docs https://pkg.go.dev/fmt
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	y = 911
	fmt.Printf("%#X\n", y)

	// printing to stdout
	fmt.Printf("Printing to stdout %#X\t%b\t%x\n", y, y, y)

	// printing to string
	s := fmt.Sprintf("Printing to string %#X\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// Fprintf, FPrintln ... -> printing to file or web server response
}
