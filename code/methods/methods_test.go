package code

import "testing"

func TestArea(t *testing.T) {
	triangle := Triangle{20.0, 10.0}

	got := triangle.Area()
	expected := 100.0

	if got != expected {
		t.Errorf("got %v expected %v", got, expected)
	}

}
