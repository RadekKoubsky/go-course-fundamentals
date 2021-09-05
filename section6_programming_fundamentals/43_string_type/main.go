package main

import (
	"fmt"
)

func main() {
	standardString := "Hello world double quotes"
	rawString := `Hello world 
					backticks`

	fmt.Println(standardString)
	fmt.Printf("%T\n", standardString)
	fmt.Println(rawString)
	fmt.Printf("%T\n", rawString)

	myString := "Hello"
	stringToBytesConversion(myString)

	printStringAsUTF8CodePoints(myString)

	printStringAsHexNumbers(myString)

	myBirthdayInBinary()
}

func myBirthdayInBinary() {
	fmt.Println("My birthday in binary")
	birthdate := "22.12.1990"
	for i := 0; i < len(birthdate); i++ {
		fmt.Printf("% 08b", birthdate[i])
	}
}

func printStringAsHexNumbers(myString string) {
	fmt.Printf("String '%v' as hex numbers:\n", myString)
	// print string as hex numbers
	for i, value := range myString {
		fmt.Printf("at index position %d we have hex %#X\n", i, value)
	}
}

func printStringAsUTF8CodePoints(myString string) {
	fmt.Printf("String '%v' as UTF-8 code points:\n", myString)
	// print string as UTF-8 code points
	for i := 0; i < len(myString); i++ {
		fmt.Printf("%#U\n", myString[i])
	}
	fmt.Println("")
}

func stringToBytesConversion(myString string) {
	fmt.Printf("String '%v' as bytes/uint8:\n", myString)
	// string value is a sequence of bytes, slices of bytes
	// we can use conversion []byte(string) to convert string value to byte array
	// byte is alias for uint8
	slicesOfBytes := []byte(myString)
	// prints byte array, each byte/uint8 represents code in ascii table, e.g. codes 72 101 108 108 111 translates
	// to H e l l o, see ascii table at https://en.wikipedia.org/wiki/ASCII
	fmt.Println(slicesOfBytes)
	fmt.Printf("%T\n", slicesOfBytes)
}
