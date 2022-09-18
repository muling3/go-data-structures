package main

import (
	"fmt"
)

// node struct
type Node struct {
	Left  *Node
	Data  int
	Right *Node
}

// Inserting into the binary search tree:: Iteratively
func (n *Node) InsertIteratively(value int) {

	if n.Data > value {

		//check if the current left is occupied
		if n.Left == nil {
			n.Left = &Node{Data: value}
		} else {
			n.Left.InsertIteratively(value)
		}

	} else if n.Data < value {

		//check if the current right is occupied
		if n.Right == nil {
			n.Right = &Node{Data: value}
		} else {
			n.Right.InsertIteratively(value)
		}
	}

}

//TODO: Insert into a binary search tree Recursively

// search a value from the binary search tree
func (n *Node) Search(value int) bool {

	//check if its greater or smaller than the root node
	if n == nil {
		return false
	}

	if n.Data == value {
		return true
	}

	if n.Data > value {
		return n.Left.Search(value)
	}

	if n.Data < value {
		return n.Right.Search(value)
	}

	return false
}

// Binary search tree traversals implemented recursively
// Preorder binary search tree traversal
func (n *Node) Preorder(root *Node) {
	count := 0
	count++
	if root == nil {
		return
	}

	//Print the data
	fmt.Print(root.Data, " ")

	root.Preorder(root.Left)
	root.Preorder(root.Right)
}

// Inorder binary search tree traversal
func (n *Node) Inorder(root *Node) {
	if root == nil {
		return
	}

	root.Inorder(root.Left)

	//Print the data
	fmt.Print(root.Data, " ")

	root.Inorder(root.Right)
}

// Postorder binary search tree traversal
func (n *Node) Postorder(root *Node) {
	if root == nil {
		return
	}

	root.Postorder(root.Left)
	root.Postorder(root.Right)

	//Print the data
	fmt.Print(root.Data, " ")
}

// finding the height of a binary search tree: Iterative way
func (n *Node) FindHeight(root *Node) int {
	countLeft := -1
	countRight := -1
	temp := root

	for temp != nil {
		countLeft++
		temp = temp.Right
	}

	for root != nil {
		countRight++
		root = root.Left
	}

	return max(countLeft, countRight)
}

// find Height Recursively
func (n *Node) FindHeightRecursively(root *Node) int {
	if root == nil {
		return -1
	}

	leftHeight := root.FindHeightRecursively(root.Left)
	rightHeight := root.FindHeightRecursively(root.Right)

	return max(leftHeight, rightHeight) + 1
}

// Deleting a leaf node from the binary search tree
func (n *Node) DeleteLeafNode(root *Node, value int) int {
	//check if its greater or smaller than the root node
	if root.Left.Data == value {
		root.Left = nil
		return 1
	}

	if root.Right.Data == value {
		root.Right = nil
		return 1
	}

	if root.Data > value {
		return root.DeleteLeafNode(root.Left, value)
	}

	return root.DeleteLeafNode(root.Right, value)
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

func main() {
	root := &Node{Data: 100}

	root.InsertIteratively(50)
	root.InsertIteratively(120)
	root.InsertIteratively(20)
	root.InsertIteratively(30)
	root.InsertIteratively(76)
	root.InsertIteratively(86)
	root.InsertIteratively(10)
	root.InsertIteratively(5)

	// check := root.Search(5)
	// if check {
	// 	fmt.Println("Found")
	// }else{
	// 	fmt.Println("Not Found")
	// }

	root.Preorder(root)
	fmt.Println()
	height := root.FindHeight(root)
	fmt.Println(height)

	// root.DeleteLeafNode(root, 120)
	// root.Preorder(root)
}
