package code

import "testing"

//three component loop
func TestFor(t *testing.T) {

	got := Repeat_for("t")
	expected := "ttttt"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}

}

// while loop
func TestWhile(t *testing.T) {
	got := Repeat_while("t")
	expected := "ttttt"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

// for each loop
func TestForeach(t *testing.T) {
	got := Repeat_foreach("ttttt")
	expected := "0t1t2t3t4t"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
