package main

import "fmt"

/*
Key understanding of channels:
CHANNELS BLOCK

Channels spec https://go.dev/ref/spec#Channel_types

Effective go https://go.dev/doc/effective_go
"Do not communicate by sharing memory; instead, share memory by communicating."

GOVERBS -> quotes from Rob Pike
https://go-proverbs.github.io/
*/
func main() {

	//blockingChannel()
	nonBlockingChannel()

	/*
	Bill Kennedy (Ultimate Go) recommends avoiding buffered channels
	as it can eventually block, thus it requires additional logic to handle full buffer

	Instead, use unbuffered channels and design your code in a way that when a value is put on the channel
	it's guaranteed that it will be read from the channel as well
	*/
	nonBlockingBufferedChannel()
	//blockingFullBufferedChannel()

}

func blockingFullBufferedChannel() {
	channel := make(chan int, 1)

	channel <- 42

	/*
		This call will block because the buffered channel is full.
	*/
	channel <- 43

	// takes a value off the channel
	fmt.Println(<-channel)
}

func nonBlockingBufferedChannel() {
	/*
		We can create a buffered channel that does not block
		until the buffer is full, we specify the buffer by passing
		size to make method.
	*/
	channel := make(chan int, 1)

	channel <- 42

	// takes a value off the channel
	fmt.Println(<-channel)
}

func nonBlockingChannel() {
	channel := make(chan int)

	/*
		By putting blocking channel operation to a separate goroutine,
		the main goroutine unblocks and reads the value from the channel
		later
	*/
	go func() {
		// puts a value to channel and blocks
		channel <- 42
	}()

	// takes a value off the channel
	fmt.Println(<-channel)
}

func blockingChannel() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	// sending a value to channel cause the channel to block
	// until somebody else receives the value from the channel
	// this call will block forever
	channel <- 42

	fmt.Println(<-channel)
}
