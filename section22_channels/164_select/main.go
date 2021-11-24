package main

import "fmt"

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan int)

	//send
	go send(evenChannel, oddChannel, quitChannel)

	//send
	receive(evenChannel, oddChannel, quitChannel)

	fmt.Println("about to exit")
}

func send(evenChannel, oddChannel, quitChannel chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenChannel <- i
		} else {
			oddChannel <- i
		}
	}
	quitChannel <- 0
}

func receive(evenChannel, oddChannel, quitChannel <-chan int) {
	for {
		/*
		Select statement allows pulling values off multiple channels
		*/
		select {
		case v := <-evenChannel:
			fmt.Println("from the even channel: ", v)
		case v := <-oddChannel:
			fmt.Println("from the odd channel: ", v)
		case v := <-quitChannel:
			fmt.Println("from the quit channel: ", v)
			return
		}
	}
}
