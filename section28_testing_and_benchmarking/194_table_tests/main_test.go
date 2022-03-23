/*
Tests must
 - be in a file that ends with “_test.go”
 - put the file in the same package as the one being tested
 - be in a func with an signature “func TestXxx(*testing.T)”
*/
package main

import "testing"

func TestWhenNumbersPassedToSum_thenReturnSumOfNumbers(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	testTable := []test{
		test{[]int{21, 21}, 42},
		test{[]int{0, 0, 0}, 0},
		test{[]int{-1, -3}, -4},
	}

	for _, testRow := range testTable {
		actual := mySum(testRow.data...)
		if actual != testRow.answer {
			t.Error("Expected", testRow.answer, "Got", actual)
		}
	}
}
