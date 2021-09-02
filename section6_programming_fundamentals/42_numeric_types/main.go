package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

// stores only values -128; 127
var signedInt int8 = -128

// stores only values 0; 255
var unsignedInt uint8 = 255

func main() {
	x = 42
	y = 42.343532
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(signedInt)
	fmt.Println(unsignedInt)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
