package main

import "fmt"

/*
Lesson from this section:
	(Almost) Always, always check your errors
	Try to avoid using underscores for swallowing errors

	Exception to this rule is situation when checking errors would cause
	infinite loop check error -> return error -> check error e.g.:
	check error of fmt.Println -> print error -> check error of fmt.Println
*/
func main() {
	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}
	// prints 6 as Println includes New Line character
	fmt.Println(n)
}
