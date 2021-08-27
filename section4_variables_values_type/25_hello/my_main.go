package main

import (
    "fmt"
    "rkoubsky.com/hello/section4_variables_values_type/25_hello/foo"
)

func main() {
    n, err := fmt.Println(foo.Foo(), "my age is", 30)
    fmt.Println("bytes written by fmt:", n, ", errors:", err)

    fmt.Printf("My name is %s and age %d\n", "radek", 30)

    // ignoring error return value
    numOfBytesWritten, _ := fmt.Println("I am using only returned bytes, but ignoring returned errors var")
    fmt.Println("Bytes written", numOfBytesWritten)
}
