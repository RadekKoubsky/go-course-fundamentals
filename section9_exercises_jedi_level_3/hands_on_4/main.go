package main

import "fmt"

func main() {
	birthdate := 1989
	for {
		if birthdate > 2021 {
			break
		}
		fmt.Println(birthdate)
		birthdate++
	}
}
