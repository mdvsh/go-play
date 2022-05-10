package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 42
	got := Add(69, -27)

	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}

}

func ExampleAdd() {
	sum := Add(69, -27)
	fmt.Println(sum)
	// Output: 42
}
