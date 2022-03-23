/*
Tests must
 - be in a file that ends with “_test.go”
 - put the file in the same package as the one being tested
 - be in a func with an signature “func TestXxx(*testing.T)”
*/
package main

import "testing"

func TestWhenNumbersPassedToSum_thenReturnSumOfNumbers(t *testing.T) {
	expected := 5
	actual := mySum(2, 3)
	if actual != expected {
		t.Error("Expected", expected, "Got", actual)
	}
}
