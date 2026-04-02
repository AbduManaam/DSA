package heap

import "fmt"

type MinHeap struct {
	arr []int
}

func (n *MinHeap) Insert(v int) {

	n.arr = append(n.arr, v)
	n.HeapifyUp(len(n.arr) - 1)
}

func (n *MinHeap) HeapifyUp(index int) {

	for index > 0 {
		parent := (index - 1) / 2
		if n.arr[parent] > n.arr[index] {
			n.arr[parent], n.arr[index] = n.arr[index], n.arr[parent]
			index = parent
		} else {
			break
		}
	}
}
func (n *MinHeap) Extract() int {

	if len(n.arr) == 0 {
		return -1
	}

	root := n.arr[0]
	last := len(n.arr) - 1
	n.arr[0] = n.arr[last]
	n.arr = n.arr[:last]
	n.HeapifyDown(0)
	return root
}
func (n *MinHeap) HeapifyDown(Index int) {

	lastIndex := len(n.arr) - 1

	for {
		left := 2*Index + 1
		right := 2*Index + 2
		smallest := Index
		if left <= lastIndex && n.arr[left] < n.arr[smallest] {
			smallest = left
		}
		if right <= lastIndex && n.arr[right] < n.arr[smallest] {
			smallest = right
		}
		if smallest != Index {
			n.arr[smallest], n.arr[Index] = n.arr[Index], n.arr[smallest]
			Index = smallest
		} else {
			break
		}
	}

}

func (n *MinHeap) Display() {

	for i := 0; i < len(n.arr); i++ {
		fmt.Print(n.arr[i]," ")
	}

}



func main() {
	h := &MinHeap{}

	// Insert elements
	s := []int{4, 5, 1, 2, 8, 9, 6, 3, 7}

	for _, ch := range s {
		h.Insert(ch)
	}

	// Display heap
	fmt.Print("Heap: ")
	h.Display()
	fmt.Println()

	// Extract elements (min first)
	fmt.Println("Extracted:", h.Extract())
	fmt.Println("Extracted:", h.Extract())

	// Display after extraction
	fmt.Print("Heap after extraction: ")
	h.Display()
	fmt.Println()
}

