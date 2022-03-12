package code

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {

	return 0.5 * t.Height * t.Width
}
