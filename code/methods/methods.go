package code

type Triangle struct {
	Width  float64
	Height float64
}

func (triangle Triangle) Area() float64 {

	return 0.5 * triangle.Height * triangle.Width
}
