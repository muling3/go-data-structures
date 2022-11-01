package queue

import "fmt"

//Implementing Queue

type Queue struct {
	items []int
}

// Enqueue operation
func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

// Dequeue operation
func (q *Queue) Dequeue() int {
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

// print queue values
func (q Queue) PrintValues() {

	for i := 0; i < len(q.items); i++ {
		fmt.Print(q.items[i], " ")
	}

	fmt.Println()
}
