package main

import "fmt"

func main() {
	channel := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Sending value to channel")
			channel <- i
		}
		/*
		The for range loop bellow reads from the channel until the channel is closed.

		We need to close the channel so that the for range loop
		does not block when the sending loop has ended.
		*/
		close(channel)
	}()

	// receive
	for v := range channel {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
