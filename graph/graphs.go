package graph

import (
	"fmt"
)

//UnDirected Graph

func UndirectedG() {
	n := 5
	graph := make([][]int, n)

	// Add edges
	addEdge := func(u, v int) {
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(1, 3)
	addEdge(3, 4)

	fmt.Println("UnDirected Graph:= ",graph)
}

func RemoveG(graph [][]int,u,v int){

   neighbour:=graph[u]
   for i,val:=range neighbour{
	 if v==val{
		graph[u] = append(neighbour[:i], neighbour[i+1:]...)
			break
	 }
   }
   neighbour=graph[v]
   for i,val:=range neighbour{
	 if u==val{
		graph[v] = append(neighbour[:i], neighbour[i+1:]...)
			break
	 }
   }

}

func IsEdgeExist(g [][]int,u,v int)bool{

	for _,val:=range g[u]{
		if val==v{
			return true
		}
	}
	return false
}

func AllNeighbors(g [][]int,u,v int){

	neighbors:=g[u]
	fmt.Print(neighbors)
}

//Directed Graph

func DirectedG(){

   n:=4
   graph:=make([][]int,n)

   addEdge:= func(u,v int){

	graph[u] = append(graph[u],v )
   }

   addEdge(0,1)
   addEdge(1,2)
   addEdge(2,3)
   addEdge(3,4)
   fmt.Println("Directed Graph:= ",graph)

}

func Bfs(g [][]int,s int){

   visited:=make(map[int]bool)
   q:=[]int{s}

   visited[s]=true

   for len(q)>0{
	node:=q[0]
	q=q[1:]

	fmt.Println("Node in BFS = ",node)

	for _,nei:=range g[node]{
		if !visited[nei]{
			visited[nei]=true

			q=append(q,nei)
		}
	}
   }

}

// Weighted Graph
type Edge struct{
	To int
	Weight int
}

func WeightedG(){

	n:=3
	graph:=make([][]Edge,n)

	addEdges:= func(u,v,w int){
		graph[u]=append(graph[u],Edge{v,w})
	}
	addEdges(0,1,10)
	addEdges(0,2,20)
	addEdges(1,3,20)
	addEdges(2,3,50)
	addEdges(2,4,40)

	fmt.Println("Weighted Graph:= ",graph)
}

