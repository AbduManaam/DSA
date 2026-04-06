package graph

import "fmt"

func UndirectedG() {
	n := 5
	graph := make([][]int, n)

	// Add edges
	addEdge := func(u, v int) {
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u) // undirected
	}

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(1, 3)
	addEdge(3, 4)

	fmt.Println(graph)
}