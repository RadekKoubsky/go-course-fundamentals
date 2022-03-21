package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	/*
		When receiving a value from a channel, we can return a status which says true for existing element in channel,
		otherwise false
	*/
	v, status := <-c

	fmt.Println(v, status)

	/*
		The last value has already been received. Returns 0 and false status

		After the last value has been received from a closed channel c, any receive from c will succeed without blocking,
		returning the zero value for the channel element and false status
	*/
	v, status = <-c

	fmt.Println(v, status)
}
