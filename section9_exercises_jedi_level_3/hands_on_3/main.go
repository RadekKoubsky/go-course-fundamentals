package main

import "fmt"

func main() {
	birthdate := 1989
	for birthdate <= 2021 {
		fmt.Println(birthdate)
		birthdate++
	}
}