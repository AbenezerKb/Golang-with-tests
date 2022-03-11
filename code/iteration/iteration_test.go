package code

import "testing"

func TestIteration(t *testing.T) {
	got := Repeat("t")
	expected := "ttttt"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
