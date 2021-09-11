package main

import "fmt"

func main() {
	switch  {
	case 0 == 0:
		fmt.Println("should print")
	case 0 == 1:
		fmt.Println("should'nt print")
	default:
		fmt.Println("this is default")




	}
}
