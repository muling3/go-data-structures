package stack

import (
	"fmt"
)

// Stack structure
type Stack struct {
	items []int
}

// push operation
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// pop operation
func (s *Stack) Pop() int {
	topElement := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return topElement
}

// pop operation
func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

// print the values
func (s Stack) PrintValues() {

	for i := 0; i < len(s.items); i++ {
		fmt.Print(s.items[i], " ")
	}

	fmt.Println()
}
