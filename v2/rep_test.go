package rep

import "testing"

func TestRep(t *testing.T) {
	rep := Repeat("a")
	expected := "aaaaa"

	if rep != expected {
		t.Errorf("expected '%s' but got '%s'", expected, rep)
	}
}

// go test -bench=.
func BenchmarkRep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
