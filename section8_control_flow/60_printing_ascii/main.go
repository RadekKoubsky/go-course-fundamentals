package main

import "fmt"

func main() {
	for i := 33; i <= 122 ; i++ {
		fmt.Printf("Ascii code %d is %#U\n", i, i)
	}
}
