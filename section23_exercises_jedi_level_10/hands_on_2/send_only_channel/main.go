package main

import "fmt"

func main() {
	/*
		original code:
		cs := make(chan<- int)
		we cannot receive a value from send only channel

		Solution: Use bidirectional channel type
	*/
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
