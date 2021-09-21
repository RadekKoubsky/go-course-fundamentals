package main

import "fmt"

func main() {
	myMap := map[string]int{
		"James": 32,
		"John":  35,
	}
	fmt.Println(myMap)
	fmt.Println(myMap["James"])
	fmt.Println(myMap["Bill"])

	// "comma ok" idiom, "ok" identifier stores the check "value exists"
	value, ok := myMap["Bill"]
	fmt.Println(value)
	fmt.Println(ok)

	if value, ok := myMap["James"]; ok {
		fmt.Printf("Value for key %s exists and is %v", "James", value)
	}
}
