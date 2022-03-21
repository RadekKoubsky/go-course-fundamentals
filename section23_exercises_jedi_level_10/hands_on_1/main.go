package main

import "fmt"

/**
get this code working:
	using func literal, aka, anonymous self-executing func
	a buffered channel
*/

func main() {
	//deadLock()
	avoidDeadlockWithFuncInSeparateGoroutine()
	avoidDeadlockWithBufferedChannel()
}

/*
	Bill Kennedy (Ultimate Go) recommends avoiding buffered channels
	as it can eventually block, thus it requires additional logic to handle full buffer

	Instead, use unbuffered channels and design your code in a way that when a value is put on the channel
	it's guaranteed that it will be read from the channel as well
*/
func avoidDeadlockWithBufferedChannel() {
	// creates buffered channel with size = 1, we can send one value to channel without blocking
	c := make(chan int, 1)

	// this will not block as the buffer (size = 1) is not full
	c <- 42

	fmt.Println(<-c)
}

func avoidDeadlockWithFuncInSeparateGoroutine() {
	c := make(chan int)

	/*
		By putting blocking channel operation to a separate goroutine,
		the main goroutine unblocks and reads the value from the channel
		later
	*/
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func deadLock() {
	c := make(chan int)

	// this will block -> deadlock
	c <- 42

	fmt.Println(<-c)
}
