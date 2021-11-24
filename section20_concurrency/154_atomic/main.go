package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
DETECTING DATA RACE in golang:
go run -race main.go
*/
func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var counter int64

	const goRoutines = 100
	var wg sync.WaitGroup
	wg.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		// executing anonymous function
		go func() {
			//time.Sleep(time.Second * 2) -> similar to runtime.Gosched()

			/*
				Gosched yields the processor, allowing other goroutines to run.
				It does not suspend the current goroutine, so execution resumes automatically.
			*/
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("numOfGoroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("After waiting for all counter goroutines. numOfGoroutines:", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
