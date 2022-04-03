package rep

import (
	"fmt"
	"testing"
)

/// func doc
func ExampleRep() {
	rep := Repeat("a", 6)
	fmt.Println(rep)
	// Output: aaaaaa
}

func TestRep(t *testing.T) {
	rep := Repeat("a", 6)
	expected := "aaaaaa"

	if rep != expected {
		t.Errorf("expected '%s' but got '%s'", expected, rep)
	}
}

// go test -bench=.
func BenchmarkRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
