package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []Person based on
// the firstname field.
type ByName[]Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

// https://pkg.go.dev/sort
func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Unsorted:")
	fmt.Println(people)
	fmt.Println("Sorted by age:")
	sort.Sort(ByAge(people))
	fmt.Println(people)
	fmt.Println("Sorted by name:")
	sort.Sort(ByName(people))
	fmt.Println(people)

}
