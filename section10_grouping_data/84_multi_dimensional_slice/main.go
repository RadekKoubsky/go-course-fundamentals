package main

import "fmt"

func main() {
		jamesBond := []string{"James", "Bond", "chocolate", "martiny"}
		fmt.Println(jamesBond)
		johnWayne := []string{"John", "Wayne", "vanilla", "whisky"}
		fmt.Println(johnWayne)

		// multi dimensional slice -> list of lists
		sliceOfSlices := [][]string{jamesBond, johnWayne}
		fmt.Println(sliceOfSlices)
}
