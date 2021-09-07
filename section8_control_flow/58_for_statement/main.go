package main

import "fmt"

/*
Effective go on for statements https://golang.org/doc/effective_go#for:

The Go for loop is similar to—but not the same as—C's.
It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

*/
func main() {
	x := 1
	// go does not have "while statement", it is replaced by "for statement with single condition" defined in language specification
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

	// we can also create for statement with no condition which is like while(true) from other languages
	y := 0
	for {
		if y > 11 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("done")

}
