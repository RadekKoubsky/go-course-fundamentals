package main

import "fmt"

func main() {
	myMap := map[string]int{
		"James": 32,
		"John":  35,
	}
	fmt.Println(myMap)

	// deletes element with specified key "James"
	fmt.Println("Deleting entry with key 'James'")
	delete(myMap, "James")
	fmt.Println(myMap)

	// If m is nil or there is no such element, delete is a no-op.
	delete(myMap, "Non-existent key")
	fmt.Println(myMap)

	if _, ok := myMap["John"]; ok {
		fmt.Println("The key 'John' is present in the map!")
	}
}
