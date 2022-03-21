package main

import "fmt"

func main() {
	/*
		original code:
		cr := make(<-chan int)
		we cannot send a value to receive only channel

		Solution: Use bidirectional channel type
	*/
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}
