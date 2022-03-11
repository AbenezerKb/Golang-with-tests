package code

import "testing"

func TestArea(t *testing.T) {

	parameterTests := []struct {
		shape Shape
		want  float64
	}{
		{Circle{0.95493}, 3},
		{Rectangle{2, 4}, 6},
	}

	for _, tdd := range parameterTests {

		got := tdd.shape.parameter()
		if got != tdd.want {
			t.Errorf("got %g want %g", got, tdd.want)
		}
	}

}
