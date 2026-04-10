package algorithms

import (
	"container/heap"
)

type Edge struct {
	to     int
	weight int
}

// Min Heap Node
type Node struct {
	vertex int
	dist   int
}

// Priority Queue (Min Heap)
type PriorityQueue []Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[:n-1]
	return node
}

// DIJKSTRA FUNCTION
func dijkstra(graph [][]Edge, src int) []int {
	n := len(graph)

	// Distance array
	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1e9 // infinity
	}
	dist[src] = 0

	// Min heap
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{vertex: src, dist: 0})

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Node)

		u := curr.vertex
		d := curr.dist

		// Skip outdated entries
		if d > dist[u] {
			continue
		}

		// Explore neighbors
		for _, edge := range graph[u] {
			v := edge.to
			weight := edge.weight

			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, Node{vertex: v, dist: dist[v]})
			}
		}
	}

	return dist
}




// func main() {
// 	graph := [][]Edge{
// 		{{1, 4}, {2, 1}}, // 0
// 		{{3, 1}},         // 1
// 		{{1, 2}, {3, 5}}, // 2
// 		{},               // 3
// 	}

// 	result := dijkstra(graph, 0)
// 	fmt.Println(result)
// }



















