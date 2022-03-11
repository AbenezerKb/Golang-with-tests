package code

import "strconv"

//three component for loop
func Repeat_for(s string) string {
	var r string

	for i := 0; i < 5; i++ {
		r += s
	}

	return r
}

//while loop
func Repeat_while(s string) string {

	var r string
	n := 0
	for n < 5 {
		r += s
		n++
	}
	return r
}

// for each loop
func Repeat_foreach(s string) string {
	var r string
	for i, v := range s {
		r = r + strconv.Itoa(i) + string(v)
	}
	return r
}
