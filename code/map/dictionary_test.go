package code

import "testing"

func TestMap(t *testing.T) {
	dictionary := map[string]string{"test": "this is the test"}
	got := search(dictionary, "test")
	want := "this is the test"

	if got != want {

		t.Errorf("got %v want %v", got, want)
	}
}
