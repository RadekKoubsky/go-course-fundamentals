package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (nError norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", nError.lat, nError.long, nError.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errMsg := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", errMsg}
	}
	return 42, nil
}

// see use of structs with error type in standard library:
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
