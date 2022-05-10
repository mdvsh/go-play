package multi

import "testing"

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
