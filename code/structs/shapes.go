package code

type Square struct {
	Height float64
}

func Perimeter(square Square) float64 {

	return 4 * square.Height
}
