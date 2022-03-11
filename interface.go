package main

import "fmt"

type AnimalSound interface {
	shout() string
}

type Dog struct {
}

type Lion struct {
}

func (d Dog) shout() string {
	return "Dogs Bark"
}

func (l Lion) shout() string {
	return "Lions roar"
}

func main() {
	animals := []AnimalSound{Dog{}, Lion{}}

	for _, animal := range animals {
		fmt.Println(animal.shout())
	}

}
