package main

import (
	"dsa/binary-tree"
	BST "dsa/binary-searchT"
	myheap "dsa/heap"
	"container/heap"
	"dsa/sorting"
	"dsa/hash-map"
    "dsa/graph"
    "dsa/adv_sortings"
	"fmt"
	
)


func main() {

	//Binary tree
	preOrder := []int{1, 2, 4, 5, 3}
	inOrder := []int{4, 2, 5, 1, 3}
    
	root:=binarytree.BuildTree(preOrder,inOrder)

	binarytree.InOrder(root)
	fmt.Println()
	binarytree.PreOrder(root)
	fmt.Println()
	binarytree.LevelPrint(root)
	fmt.Println()


	f:=binarytree.Insert(root,100)
	fmt.Print("After Insert: ")
	binarytree.LevelPrint(f)
	fmt.Println()

	fmt.Println("Search=",binarytree.Searchh(root,150))
	fmt.Println("Search=",binarytree.Searchh(root,100))

	fmt.Println()

	fmt.Println("Updated: ",binarytree.Update(root,4, 40 ))

	r:=binarytree.Delete(root,3 )
	fmt.Print("After Delete: ")
	binarytree.LevelPrint(r)
	fmt.Println()

	fmt.Println("Pyramid:")

	binarytree.LevelPrintPyramid(root)
	fmt.Println()
	
	fmt.Println("Search=",binarytree.Search(root,15))

	//Binary-Search Tree

	fmt.Println("Binary Search Tree:")
	var BSTroot *BST.Node

	arr2 := []int{1, 14, 5, 82, 1, 6, 5}
	for _, v := range arr2 {
		BSTroot = BST.Insert(BSTroot, v)
	}
	BST.Inorder(BSTroot)


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

   //Insertion Sort
    arr := []int{1, 2, 4, 5, 8, 7, 3}
	fmt.Println("Insertion sorted:",sorting.Insertion(arr))

   //Selection Sort

   arr1 := []int{1, 2, 4, 5, 8, 7, 3}
	fmt.Println("Selection sorted:",sorting.Insertion(arr1))

//---------------------------------------------

    fmt.Println()

	p := &myheap.Pq{
		{Val: "T1", Priority: 1},
		{Val: "T2", Priority: 5},
		{Val: "T3", Priority: 7},
		{Val: "T4", Priority: 2},
		{Val: "T5", Priority: 3},
	}
	heap.Init(p)
	heap.Push(p,&myheap.MaxHeap{Val: "T12",Priority: 10})
	heap.Push(p,&myheap.MaxHeap{Val:"T34",Priority: 11})

	fmt.Println("Priority Queue (max heap order):")
	fmt.Println()
    for p.Len()>0{
		item:=heap.Pop(p).(*myheap.MaxHeap)
		fmt.Printf("Value:%s,Priority:%d\n",item.Val,item.Priority)
	}
	
	fmt.Println("Height of tree:", binarytree.Height(root))
	fmt.Println("Total nodes:", binarytree.CountNode(root))
	fmt.Println("Leaf nodes:", binarytree.CountLeaf(root))
	fmt.Println("Internal nodes:", binarytree.CountInterval(root))

	//Hash Map

	hm := hashmap.BuildHashMap(5)

	hm.Put("apple", 10)
	hm.Put("banana", 20)
	hm.Put("grape", 30)
	hm.Put("apple", 50) // update
    fmt.Println("Hash Elements: ")
	hm.Display()

	val, found := hm.Get("apple")
	fmt.Println("apple:", val, found)

	hm.Remove("banana")
	hm.Display()

	// Graph
	fmt.Println("Graph ")
	graph.UndirectedG()
	graph.DirectedG()
	graph.WeightedG()
	graph.MUndirectedG()
	graph.MdirectedG()
	graph.MWeightedG()

	// Graph Matrix

	ms:=5
	edges:=[][]int{
		{0, 1},
		{0, 2},
		{1, 3},
		{2, 3},

	}

	fmt.Println("Undirected Matrix =",graph.BuildUndirectedMatrix(ms,edges))
	fmt.Println("Directed Matrix =",graph.BuildDirectedMatrix(ms,edges))
	fmt.Println("Weighted Matrix =",graph.BuildWeightedMatrix(ms,edges))
    graph.PrintMatrix(edges)

	// Quick Sort

	arr3:= []int{10, 7, 8, 9, 1, 5}
	advsortings.QuickSort(arr1, 0, len(arr)-1)
	fmt.Println("QuickSort Array: ",arr3)

}
