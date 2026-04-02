package heap

// import (
// 	"container/heap"
// 	"fmt"
// )

type MaxHeap struct {
	Val      string
	Priority int
}

type Pq []*MaxHeap

func (p Pq) Len() int {
	return len(p)
}

func (p Pq) Less(i, j int) bool {
	return p[i].Priority > p[j].Priority
}

func (p Pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pq) Push(Val interface{}) {
	*p = append(*p, Val.(*MaxHeap))
}

func (p *Pq) Pop() interface{} {

	old := *p
	s := len(old)
	last := old[s-1]
	*p = old[0 : s-1]
	return last

}

// func main() {

// 	p := &Pq{
// 		{Val: "T1", Priority: 1},
// 		{Val: "T2", Priority: 5},
// 		{Val: "T3", Priority: 7},
// 		{Val: "T4", Priority: 2},
// 		{Val: "T5", Priority: 3},
// 	}
// 	heap.Init(p)
// 	heap.Push(p,&MaxHeap{Val: "T12",Priority: 10})
// 	heap.Push(p,&MaxHeap{Val:"T34",Priority: 11})

// 	fmt.Println("Priority Queue (max heap order):")
//     for p.Len()>0{
// 		item:=heap.Pop(p).(*MaxHeap)
// 		fmt.Printf("Value:%s,Priority:%d\n",item.Val,item.priority)
// 	}
// }


// You need:

// heap.Pop(p).(*MaxHeap)

// because:

// 👉 heap.Pop() returns interface{}
// 👉 You must convert it back to your actual type (*MaxHeap)

