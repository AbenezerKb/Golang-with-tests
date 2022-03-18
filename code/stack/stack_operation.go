package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	value []int
	index int
}

func (s *stack) pop() int {
	temp := s.value[s.index-1]
	s.index--
	return temp
}

func (s *stack) insert(num int) {
	s.value = append(s.value, num)
	s.index++
}

func main() {

	s := stack{}
	s.insert(7)
	s.insert(6)
	s.insert(5)
	s.insert(4)

	fmt.Println(strconv.Itoa(s.pop()))
	fmt.Println(strconv.Itoa(s.pop()))
	fmt.Println(strconv.Itoa(s.pop()))
	fmt.Println(strconv.Itoa(s.pop()))

}
