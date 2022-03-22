package main

import (
	"fmt"
	"os"
)

/*
   Use the following three types of functions to log your errors:

	log.Println("err happened", err)
	log.Fatalln(err)
	log.Panicln(err)
*/
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
}
