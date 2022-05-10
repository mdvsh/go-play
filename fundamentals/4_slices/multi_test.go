package multi

import (
	"reflect"
	"testing"
)

func TestMult(t *testing.T) {
	t.Run("collection of 5 nums", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5} // pass in a slice to abide by func defn
		expected := 120
		got := Multi(nums)

		if expected != got {
			t.Errorf("expected %d, got %d, given %v", expected, got, nums)
		}
	})
	t.Run("collection of arbitrary nums", func(t *testing.T) {
		nums := []int{10, 9, 8}
		expected := 720
		got := Multi(nums)

		if expected != got {
			t.Errorf("expected %d, got %d, given %v", expected, got, nums)
		}
	})
	t.Run("empty nums slice ?", func(t *testing.T) {
		nums := []int{}
		expected := 0
		got := Multi(nums)
		if expected != got {
			t.Errorf("expected %d, got %d, given %v", expected, got, nums)
		}
	})
}

func TestMultAll(t *testing.T) {
	assertor := func(t *testing.T, expected, got []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %d, got %d", expected, got)
		}
	}

	t.Run("(same size) multiply corresponding elements", func(t *testing.T) {
		got := MultiAll([]int{4, 7, 11}, []int{6, 1, 9})
		expected := []int{24, 7, 99}
		// if got != expected {
		// 	t.Errorf("expected %d, got %d", expected, got)
		// }

		// Can't use equality operator with slice.
		// Use, DeepEqual: https://pkg.go.dev/reflect#DeepEqual
		assertor(t, expected, got)
	})
	t.Run("(diff size, left more) multiply corresponding elements", func(t *testing.T) {
		got := MultiAll([]int{4, 6, 3}, []int{1, 4})
		expected := []int{4, 24, 3}
		assertor(t, expected, got)
	})
	t.Run("(diff size, right more) multiply corresponding elements", func(t *testing.T) {
		got := MultiAll([]int{1, 4}, []int{4, 6, 3})
		expected := []int{4, 24, 3}
		assertor(t, expected, got)
	})
	t.Run("empty slice", func(t *testing.T) {
		got := MultiAll([]int{}, []int{})
		expected := []int{}
		assertor(t, expected, got)
	})
}
