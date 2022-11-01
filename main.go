package main

import (
	"fmt"
	_ "fmt"

	bst "github.com/AlexanderMuli/go-data-structures/bst"
)

func main() {
	root := &bst.Node{Data: 100}

	root.InsertIteratively(50)
	root.InsertIteratively(120)
	root.InsertIteratively(20)
	root.InsertIteratively(30)
	root.InsertIteratively(76)
	root.InsertIteratively(86)
	root.InsertIteratively(10)
	root.InsertIteratively(5)

	root.Preorder(root)
	fmt.Println()

}
