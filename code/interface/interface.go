package code

import (
	"math"
	
)

type Shape interface {
	parameter() float64
}

type Circle struct {
	raduis float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) parameter() float64 {
	return 2 * math.Pi * c.raduis
}

func (r Rectangle) parameter() float64 {
	return r.width + r.height
}


// func main() {
// 	animals := []AnimalSound{Dog{}, Lion{}}

// 	for _, animal := range animals {
// 		fmt.Println(animal.shout())
// 	}

// }
