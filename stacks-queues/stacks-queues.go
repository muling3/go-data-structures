package main

import (
	"fmt"
	"strings"
)

type Queue struct{
	items []int
}

//Stack structure
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

// print the values
func (s Stack) PrintValues() {

	for i := 0; i < len(s.items); i++ {
		fmt.Print(s.items[i], " ")
	}

	fmt.Println()
}

//Implementing Queue
//Enqueue operation
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

//Dequeue operation
func (q *Queue) Dequeue() int{
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

//print queue values
func (q Queue) PrintValues() {

	for i := 0; i < len(q.items); i++ {
		fmt.Print(q.items[i], " ")
	}

	fmt.Println()
}

func main() {
	myStack := Stack{}
	myQueue := Queue{}

	myStack.Push(12)
	myStack.Push(13)
	myStack.Push(14)
	myStack.Push(15)
	myStack.Push(16)
	top := myStack.Pop()

	// myStack.PrintValues()

	myQueue.Enqueue(12)
	myQueue.Enqueue(13)
	myQueue.Enqueue(14)
	myQueue.Enqueue(15)
	myQueue.Enqueue(16)
	queueTop := myQueue.Dequeue()

	// myQueue.PrintValues()

	fmt.Println("The Top Element at the stack is", top)
	fmt.Println("The First Element at the queue is", queueTop)
	fmt.Println(strings.Join([]string{"Jose", "Carmago", "With", "IEBC", "Register"}, " "))
}
