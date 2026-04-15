package graph

import (
	"fmt"
)

//UnDirected Graph

func paraUndirectedG(edge [][]int,s int)[][]int{

   graph:=make([][]int,s)

   for _,val:=range edge{
	
	u,v := val[0],val[1]
	graph[u]=append(graph[u],v)
	graph[v]=append(graph[v],u)

   }
   return graph

}

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


//------------------------------------------------------------------------------------

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

	if u < 0 || u >= len(g) {
		return false
	}

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



//------------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------------

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


//------------------------------------------------------------------------------------

func Dfs(graph [][]int,node int,visited map[int]bool){

   if visited[node]{
	return
   }
	visited[node]=true

	fmt.Print(node," ")
   for _,nei:=range graph[node]{
	Dfs(graph,nei,visited)
   }

   fmt.Println("Length of graph = ",len(graph))

}

// func main() {
// 	// Adjacency list graph
// 	graph := [][]int{
// 		{1, 2}, // 0
// 		{0, 3}, // 1
// 		{0},    // 2
// 		{1},    // 3
// 	}

// 	visited := make(map[int]bool)

// 	startNode := 0

// 	Dfs(graph, startNode, visited)
// }



//------------------------------------------------------------------------------------

// DFS Cycle Detect

func hashCycleDFS(graph [][]int) bool {
	visited := make([]bool, len(graph))

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

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			if dfs(i, -1) {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	graph := [][]int{
// 		{1},       // 0
// 		{0, 2},    // 1
// 		{1, 3},    // 2
// 		{2, 0},    // 3 → cycle
// 	}

// 	fmt.Println(hashCycleDFS(graph)) // true
// }

//------------------------------------------------------------------------------------



// BFS Cycle Detect

type Pair struct {
	node   int
	parent int
}

func hasCycleBFS(graph [][]int) bool {
	n := len(graph)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {          //Handle disconnected graph
		if !visited[i] {

			// start BFS from this node
			
			queue := []Pair{{i, -1}}  // Here: queue already exists,you are adding a new element (not appending concept)
			visited[i] = true

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]

				node := curr.node
				parent := curr.parent

				for _, nei := range graph[node] {

					if !visited[nei] {
						visited[nei] = true   // go forward
						queue = append(queue, Pair{nei, node})

					} else if nei != parent {						
						return true  // visited and not parent → cycle
					}
				}
			}
		}
	}
	return false
}

// func main() {
// 	graph := [][]int{
// 		{1},        // 0
// 		{0, 2},     // 1
// 		{1, 3},     // 2
// 		{2, 1},     // 3 → cycle here
// 	}

// 	fmt.Println(hasCycleBFS(graph)) // true
// }


//------------------------------------------------------------------------------------

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


//------------------------------------------------------------------------------------
