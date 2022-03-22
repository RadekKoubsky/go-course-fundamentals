package main

import "fmt"

// json package has a real world use case of panic and recover
func main() {
	// f panicked after g, but deferred function called recover
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		// two statements on one line, separated by semicolon, common idiom in golang
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	// g panicked here, g's deferred functions were called,
	// f starts panicking after g returns, deferred function in f executes recover() function, and we recover to normal flow
	g(0)
	// this code never executes as f panics after g(0) call
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
