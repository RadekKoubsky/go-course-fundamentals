package main

import "fmt"

func main() {
	x := 0
	for {
		x++
		if x > 100 {
			// stop here and also stop the entire loop
			break
		}
		if x%2 != 0 {
			// stop here and move to the start of the next iteration
			// Important note: code after "continue statement" is skipped and not executed
			continue
		}
		fmt.Println(x)
	}
}
