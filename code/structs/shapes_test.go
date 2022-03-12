package code

import "testing"

func TestPerimeter(t *testing.T) {
	square := Square{10.0}
	got := Perimeter(square)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
