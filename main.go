package main

import (
	"dsa/binary-tree"
	myheap "dsa/heap"
	"container/heap"
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
	m := &myheap.MinHeap{}
	s := []int{4, 5, 1, 2, 8, 9, 6, 3, 7}

	for _, ch := range s {
		m.Insert(ch)
	}
	fmt.Println(m.Extract())
    fmt.Println("Heap Values:")
	m.Display()
	fmt.Println()

	//Priority Queue
	 h:=&myheap.InHeap{4,1,2,8,7}
   heap.Init(h)
   heap.Push(h, 10)
   fmt.Printf("Min: %d\n",(*h)[0])
   
   copyHeap := make(myheap.InHeap, len(*h))
    copy(copyHeap, *h)

   fmt.Println("Priority Queue :")
   for h.Len()>0{
	fmt.Print(heap.Pop(h)," ")
   }

   // Shows elements in priority order
    // Original heap remains unchanged
    
	fmt.Println()
	
    fmt.Println("Priority Queue :")
    for copyHeap.Len() > 0 {
    fmt.Println(heap.Pop(&copyHeap)," ")
}


}