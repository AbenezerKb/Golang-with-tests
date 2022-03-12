package code

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got := Sum(numbers)
	expected := 55

	if got != expected {
		t.Errorf("got %v expected %v", got, expected)
	}
}
