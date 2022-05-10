package main

import "testing"

func TestMain(t *testing.T) {
	expected := "hi madhav"
	got := Hello("madhav")

	if expected != got {
		t.Errorf("got %q want %q", got, expected)
	}

}
