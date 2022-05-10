package main

import "testing"

func TestHello(t *testing.T) {
	assertor := func(t *testing.T, expected, got string) {
		t.Helper()
		// needed to tell the test suite that this method is a helper, points to right buggy line
		if expected != got {
			t.Errorf("got %q want %q", got, expected)
		}
	}
	t.Run("basic hello test", func(t *testing.T) {
		expected := "hi madhav"
		got := Hello("madhav", "")
		assertor(t, expected, got)

	})
	t.Run("call empty with no args", func(t *testing.T) {
		expected := "hi world"
		got := Hello("", "")
		assertor(t, expected, got)
	})
	t.Run("hindi hello", func(t *testing.T) {
		expected := "namaste madhav"
		got := Hello("madhav", "hindi")
		assertor(t, expected, got)
	})
	t.Run("french hello", func(t *testing.T) {
		expected := "bonjour madhav"
		got := Hello("madhav", "french")
		assertor(t, expected, got)
	})
}
