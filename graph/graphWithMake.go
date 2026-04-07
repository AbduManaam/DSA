package graph

import (
	"fmt"
)

//------------------------------------------------------------------------------------

func MUndirectedG() {

	graph:=make(map[int][]int)

	addEdge:= func(u,v int){
		graph[u]=append(graph[u],v)
		graph[v]=append(graph[v],u)
	}

	addEdge(0,1)
	addEdge(1,0)
	addEdge(0,3)
	addEdge(3,0)
	addEdge(0,2)

	fmt.Println("UnDirected Graph (make):= ",graph)
}

//------------------------------------------------------------------------------------

// Cycle Detection

func hasCycle(graph map[int][]int) bool {
	visited := make(map[int]bool)

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true

		for _, nei := range graph[node] {
			if !visited[nei] {
				if dfs(nei, node) {
					return true
				}
			} else if nei != parent {
				return true 
			}
		}
		return false
	}

	for node := range graph {
		if !visited[node] {
			if dfs(node, -1) {
				return true
			}
		}
	}

	return false
}

//------------------------------------------------------------------------------------

func main() {
	graph := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1, 3},
		3: {2, 0}, // cycle here
	}

	fmt.Println(hasCycle(graph)) // true
}

//------------------------------------------------------------------------------------

func MdirectedG() {

	graph:=make(map[int][]int)

	addEdge:= func(u,v int){
		graph[u]=append(graph[u],v)
	}

	addEdge(0,1)
	addEdge(1,2)
	addEdge(2,3)

	fmt.Println("Directed Graph (make):= ",graph)
}

//------------------------------------------------------------------------------------

type Edges struct{
	To int
	Weight int
}

func MWeightedG(){

	graph:=make(map[int][]Edges)

	addEdges:= func(u,v,w int){
		graph[u]=append(graph[u],Edges{v,w})
	}
	addEdges(0,1,10)
	addEdges(0,2,20)
	addEdges(1,3,20)
	addEdges(2,3,50)
	addEdges(2,4,40)

	fmt.Println("Weighted Graph(make):= ",graph)
}