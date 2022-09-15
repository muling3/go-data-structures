package main

import (
	"fmt"
	"strings"
)

// Node structure
type node struct {
	data int
	next *node
}

// LinkedList structure
type LinkedList struct {
	head   *node
	length int
}

// Operations
func (l *LinkedList) AddBeg(n *node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

// add at the end
func (head *node) AddEnd(n *node) *node {
	if head.data == 0 {
		head.data = n.data
		return head
	}

	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = n

	return head
}

// find a given value
func (l *LinkedList) findByValue(value int) string {
	temp := l.head

	for l.length != 0 {
		if temp.data == value {
			return "Present"
		}
		temp = temp.next
		l.length--
	}

	return "Not Present"
}

//iterative way
func (n *node) findByValue(head *node, value int) string {
	temp := head

	for temp.next != nil {
		if temp.data == value {
			return "Present"
		}
		temp = temp.next
	}

	if temp.data == value {
		return "Present"
	}

	return "Not Present"
}

//finding in a recursuive way
func (n *node) findByValueRecursive(head *node, value int) string {
	temp := head

	if temp.data == value {
		return "Present"
	}

	if temp.next == nil {
		return "Not Present"
	}

	return head.findByValueRecursive(temp.next, value)
}


//deleting a node by value
func (l *LinkedList) deleteByValue(value int, ll LinkedList) {

	//check if value is Present
	check := ll.findByValue(value)

	if strings.Compare(check, "Present") == 0 {
		temp := l.head
		if temp.data == value {
			l.head = temp.next
			l.length--
			return
		}

		for temp.next.data != value {
			temp = temp.next
		}

		temp.next = temp.next.next
		l.length--
	} else {
		fmt.Println("The value does not exist")
	}
}

/*
	Deletes the node at the given position
	Count starts from 1
	NOTE: A node can not have a value 0(Zero)
*/
func (n *node) deleteAtPosition(pos int, head *node) {
	temp := head
	
	//check if position is equal to 1 to change the head address
	if pos == 1 {
		temp.data = temp.next.data
		temp.next = temp.next.next
		return
	}else if pos > head.CountNodes(head){
		fmt.Println("Make sure the list has enough elements")
		return
	}else{

		for pos != 1{
			temp = temp.next
			pos--					
		}
		if temp.next != nil {
			temp.data = temp.next.data
			temp.next = temp.next.next
			return
		}

		temp.data = 0
	}

}

//Deletes a node which contains the value given as an argument
func (n *node) deleteByValue(head *node, value int) {

	temp := head
	//check if value is Present
	check := head.findByValue(head, value)

	if strings.Compare(check, "Present") == 0 {
		if temp.data == value {
			head.data = temp.next.data
			head.next = temp.next.next
			return
		}

		for temp.next.data != value{
			temp = temp.next
		}

		temp.next = temp.next.next
	} else {
		fmt.Println("The value does not exist")
	}
}

//Prints all the value to the console
func (l LinkedList) PrintValues() {
	current := l.head

	for l.length != 0 {
		fmt.Print(current.data, " ")
		current = current.next
		l.length--
	}

	fmt.Println()

}

//Prints all the value to the console
func (n node) PrintValues(head *node) {
	current := head

	for current != nil && current.data != 0 {
		fmt.Print(current.data, " ")
		current = current.next
	}

	fmt.Println()
}

//Reverses a linked list
func (n *node) ReverseList(head *node) *node{
	prev := &node{}
	next := &node{}

	for head != nil{
		next = head.next
		head.next = prev
		prev = head
		head = next
	}

	head = prev

	return head
}

//count the number of nodes in the list and then returns the count(Integer)
func (n node) CountNodes(head *node) int{
	current := head
	count := 0

	for current != nil && current.data != 0{
		count++
		current = current.next
	}

	return count

}

func main() {
	myList := LinkedList{}
	myList.AddBeg(&node{6, nil})
	myList.AddBeg(&node{12, nil})
	myList.AddBeg(&node{13, nil})

	head := &node{4, nil}

	head.AddEnd(&node{6, nil})
	head.AddEnd(&node{12, nil})
	head.AddEnd(&node{13, nil})
	head.AddEnd(&node{14, nil})
	head.AddEnd(&node{24, nil})
	head.AddEnd(&node{84, nil})
	head.AddEnd(&node{34, nil})
	head.AddEnd(&node{54, nil})
	head.AddEnd(&node{94, nil})

	// head.PrintValues(head)

	fmt.Println("Before removing")
	head.PrintValues(head)
	new_head := head.ReverseList(head)
	fmt.Println("After Reversing")
	head.PrintValues(new_head)

	// fmt.Println(head.findByValueRecursive(head, 6))
	// myList.deleteByValue(13, myList)
	// myList.PrintValues()
	// fmt.Println(myList.findByValue(6))

}
