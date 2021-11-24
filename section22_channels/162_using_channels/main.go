package main

import "fmt"

func main() {
	channel := make(chan int)


	// send
	go foo(channel)

	// receive
	bar(channel)

	fmt.Println("about to exit")
}

/*
Param is send only channel

The scope of param is only this function, thus we can
convert it to send only channel used only in this function's scope
*/
func foo(channel chan<- int) {
	channel <- 42
}

// receive
func bar(channel <-chan int) {
	fmt.Println(<-channel)
}
