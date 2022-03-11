package code

func Repeat(s string) string {
	var r string
	for i := 0; i < 5; i++ {
		r += s
	}

	return r
}
