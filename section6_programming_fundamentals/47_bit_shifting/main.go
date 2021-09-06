package main

import (
	"fmt"
)

const (
	// ignore first iota = 0
	_ = iota
	// kbConst = 1024
	// we shift 1 by 1(iota = 1) * 10 zeros -> 10000000000
	kbConst = 1 << (iota * 10)
	// we shift 1 by 2(iota = 2) * 10 zeros -> 10000000000 0000000000
	mbConst = 1 << (iota * 10)
	// we shift 1 by 3(iota = 3) * 10 zeros -> 10000000000 0000000000 0000000000
	gbConst = 1 << (iota * 10)
)

func main() {
	bitShifting()

	printBytesWithoutIota()

	fmt.Println("Printing bytes using iota")
	fmt.Printf("%d\t\t\t%b\n", kbConst, kbConst)
	fmt.Printf("%d\t\t\t%b\n", mbConst, mbConst)
	fmt.Printf("%d\t\t%b\n", gbConst, gbConst)
}

func printBytesWithoutIota() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	// we can see that each time we are shifting over 10 zeros, this is good candidate to use iota
	fmt.Println("Printing bytes without iota")
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

func bitShifting() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	z := x >> 1
	fmt.Printf("%d\t\t%b\n", z, z)
}
