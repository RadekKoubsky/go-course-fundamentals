package saying

import (
	"fmt"
	"testing"
)

/*
Test coverage

We can use the “-cover” flag to run coverage analysis on our code.

go test -cover

We can use the flag and required file name “-coverprofile <some file name>” to write our coverage analysis to a file.
code:

go test -coverprofile=c.out

show in browser (visual for each line):
go tool cover -html=c.out

*/
func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear James" {
		t.Error("got", s, "expected", "Welcome my dear James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
