package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**

in addition to the main goroutine, launch two additional goroutines
	each additional goroutine should print something out

use waitgroups to make sure each goroutine finishes before your program exists


*/

var wg sync.WaitGroup

func main() {
	fmt.Println("begin CPUs\t\t", runtime.NumCPU())
	fmt.Println("begin Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("in goroutine #1")
		wg.Done()
	}()

	go func() {
		fmt.Println("in goroutine #2")
		wg.Done()
	}()
	/*
	We need to wait for all goroutines in this waitgroup to complete in order to
	see Println statements called in the goroutines.

	Without waiting for the goroutines, the main goroutine could possibly
	exit before the other goroutines, thus we would not see Println statements
	called in the goroutines every time.
	*/

	fmt.Println("mid CPUs\t", runtime.NumCPU())
	fmt.Println("mid Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("end CPUs\t", runtime.NumCPU())
	fmt.Println("end Goroutines\t", runtime.NumGoroutine())
}
