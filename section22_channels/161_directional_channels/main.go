package main

import "fmt"

/*
Bidirectional channels: you send to it and receive from it

Channel with one direction:
you can only send to channel
or
you can only receive from channel

*/
func main() {
	bidirectionalChannel()
	//receiveOperationOnSendOnlyChannel()
	//sendOperationOnReceiveOnlyChannel()

	// assigning channels of different types
	// generalChannel = specificChannel -> does not work
	// specificChannel = generalChannel -> does work, possible to do conversion
}

func receiveOperationOnSendOnlyChannel() {
	channel := make(chan<- int, 2)

	channel <- 42
	channel <- 43

	/*
		Does not compile cannot receive from send-only channel
		fmt.Println(<-channel)
		fmt.Println(<-channel)*/
	fmt.Println("------")
	fmt.Printf("%T\n", channel)
}

func sendOperationOnReceiveOnlyChannel() {
	channel := make(<-chan int, 2)

	/*
		Does not compile, send to the receive-only type <-chan int
		channel <- 42
		channel <- 43*/

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("------")
	fmt.Printf("%T\n", channel)
}

func bidirectionalChannel() {
	channel := make(chan int, 2)

	channel <- 42
	channel <- 43

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("------")
	fmt.Printf("%T\n", channel)
}
