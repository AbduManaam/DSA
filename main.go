package main

import (
	"dsa/binary-tree"
	"dsa/heap"
	"fmt"
)


func main() {
	preOrder := []int{1, 2, 4, 5, 3}
	inOrder := []int{4, 2, 5, 1, 3}
    
	root:=binarytree.BuildTree(preOrder,inOrder)

	binarytree.InOrder(root)
	fmt.Println()
	binarytree.PreOrder(root)
	fmt.Println()
	binarytree.LevelPrint(root)
	fmt.Println()

	fmt.Println("Pyramid:")

	binarytree.LevelPrintPyramid(root)
	fmt.Println()
	
	fmt.Println("Search=",binarytree.Search(root,15))

    //HEAP
	m := &heap.MinHeap{}
	s := []int{4, 5, 1, 2, 8, 9, 6, 3, 7}

	for _, ch := range s {
		m.Insert(ch)
	}
	fmt.Println(m.Extract())
    fmt.Println("Heap Values:")
	m.Display()
	fmt.Println()

}