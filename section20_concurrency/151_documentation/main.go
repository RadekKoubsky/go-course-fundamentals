package main

import "fmt"

func doSomething(x int) int {
	return x * 5
}

func main() {
	// if you need to access a return value from a function running in goroutine, put
	// that value to channel and you access later
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	fmt.Println(<-ch)
}

