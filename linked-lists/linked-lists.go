package linkedlist

import (
	"fmt"
)

// Node structure
type Node struct {
	Data int
	Next *Node
}

// add at the end
func (head *Node) AddEnd(n *Node) {
	if head.Data == 0 {
		head.Data = n.Data
	}

	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = n
}

// iterative way
func (n *Node) FindByValue(head *Node, value int) bool {
	temp := head

	for temp.Next != nil {
		if temp.Data == value {
			return true
		}
		temp = temp.Next
	}

	return temp.Data == value
}

// finding in a recursuive way
func (n *Node) FindByValueRecursive(head *Node, value int) bool {
	temp := head

	if temp.Data == value {
		return true
	}

	if temp.Next == nil {
		return false
	}

	return head.FindByValueRecursive(temp.Next, value)
}

/*
Deletes the node at the given position
Count starts from 1
NOTE: A node can not have a value 0(Zero)
*/
func (n *Node) DeleteAtPosition(pos int, head *Node) {
	temp := head

	//check if position is equal to 1 to change the head address
	if pos == 1 {
		temp.Data = temp.Next.Data
		temp.Next = temp.Next.Next
		return
	} else if pos > head.CountNodes(head) {
		fmt.Println("Make sure the list has enough elements")
		return
	} else {

		for pos != 1 {
			temp = temp.Next
			pos--
		}
		if temp.Next != nil {
			temp.Data = temp.Next.Data
			temp.Next = temp.Next.Next
			return
		}

		temp.Data = 0
	}
}

// Deletes a node which contains the value given as an argument
func (n *Node) DeleteByValue(head *Node, value int) bool {
	temp := head

	if head.FindByValue(head, value) {
		if temp.Data == value {
			head.Data = temp.Next.Data
			head.Next = temp.Next.Next
			return true
		}

		for temp.Next.Data != value {
			temp = temp.Next
		}

		temp.Next = temp.Next.Next
	}

	return false
}

// Prints all the value to the console
func (n *Node) PrintValues(head *Node) {
	current := head

	for current != nil && current.Data != 0 {
		fmt.Print(current.Data, " ")
		current = current.Next
	}

	fmt.Println()
}

// Reverses a linked list
func (n *Node) ReverseList(head *Node) *Node {
	var prev *Node
	var next *Node

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	head = prev

	return head
}

// count the number of nodes in the list and then returns the count(Integer)
func (n *Node) CountNodes(head *Node) int {
	current := head
	count := 0

	for current != nil && current.Data != 0 {
		count++
		current = current.Next
	}

	return count
}
