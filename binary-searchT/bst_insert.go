package binarysearcht

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}
	return root
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Println("Inorder Value: ",root.Val)
	Inorder(root.Right)
}

// func main() {
// 	var root *Node

// 	arr := []int{1, 14, 5, 82, 1, 6, 5}
// 	for _, v := range arr {
// 		root = Insert(root, v)
// 	}
// 	Inorder(root)
// }

// To use this function: (BT)

// BuildTree(preOrder, inOrder)

// 👉 The inorder must be sorted
// 👉 AND both traversals must come from a valid BST


//---------------------------------------------------------------------------------------
