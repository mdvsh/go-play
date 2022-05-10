package iter

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	t.Run("string repeat", func(t *testing.T) {
		expected := "aaaaa"
		got := Repeat("a", 5)
		if expected != got {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}

// repats a given string character specified number of times
func ExampleRepeat() {
	repeated := Repeat("m", 6)
	fmt.Println(repeated)
	// Output: mmmmmm
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
