package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
DETECTING DATA RACE in golang:
go run -race main.go
*/
func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	counter := 0

	const goRoutines = 100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		// executing anonymous function
		go func() {
			v := counter
			//time.Sleep(time.Second * 2) -> similar to runtime.Gosched()

			/*
			Gosched yields the processor, allowing other goroutines to run.
			It does not suspend the current goroutine, so execution resumes automatically.
			*/
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("numOfGoroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("After waiting for all counter goroutines. numOfGoroutines:", runtime.NumGoroutine())
	// due to race condition in read/write to shared counter var, counter will have random values
	fmt.Println("Counter: ", counter)
}
