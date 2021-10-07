package main

import "fmt"

func main() {
	// an anonymous struct, it does not have a name
	// if you only need a struct in one little area, you can define it as anonymous struct
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
