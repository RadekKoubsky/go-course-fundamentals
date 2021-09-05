package main

import "fmt"

func main() {
	myString := "H"
	fmt.Println(myString)

	byteArray := []byte(myString)
	fmt.Println(byteArray)

	myByte := byteArray[0]
	fmt.Println(myByte)
	fmt.Printf("%T\n", myByte)
	fmt.Printf("%b\n", myByte)
	fmt.Printf("%#X\n", myByte)


}
