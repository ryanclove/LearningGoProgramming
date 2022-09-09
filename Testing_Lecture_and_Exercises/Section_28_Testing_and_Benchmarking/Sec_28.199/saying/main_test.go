package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear James" {
		t.Error("Got:", s, "\tExpected:", "Welcome my dear James")
	}
}

func ExampleGreet(t *testing.T) {
	fmt.Println(Greet("James"))
	// Output:
	// Welcome my dear James
}

// Runs the code a lot of times until it gets a statistcally accurate
// measurement of how long it took to run it - give it a loop
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
