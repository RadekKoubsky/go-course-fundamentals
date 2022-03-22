package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	file, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer file.Close()

	fmt.Println("check the log.txt file in the directory")
}
