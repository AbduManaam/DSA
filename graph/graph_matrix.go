package graph

import (
	"fmt"
)

func BuildUndirectedMatrix(s int, graph [][]int) [][]int {

	matrix := make([][]int, s)

	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)  //creates one row per vertex
	}

	for _, e := range graph {
		u, v := e[0], e[1]
		matrix[u][v] = 1
		matrix[v][u] = 1
	}
	return matrix

}


func BuildDirectedMatrix(s int, graph [][]int) [][]int {

	matrix := make([][]int, s)

	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)
	}

	for _, e := range graph {
		u, v := e[0], e[1]
		matrix[u][v] = 1
	}
	return matrix

}

//-----------------------------------------------------------------------------------------------------------------------------


func BuildWeightedMatrix(s int, graph [][]int) [][]int {
	matrix := make([][]int, s)

	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)
	}

	for _, e := range graph {
		if len(e) != 3 {
			panic("Each edge must have exactly 3 values: [u, v, w]")
		}

		u, v, w := e[0], e[1], e[2]

		if u < 0 || v < 0 || u >= s || v >= s {
			panic("Invalid vertex index")
		}

		matrix[u][v] = w
		matrix[v][u] = w 
	}

	return matrix
}

// func main() {
// 	// Number of vertices
// 	size := 4

// 	// Graph edges: [u, v, weight]
// 	graph := [][]int{
// 		{0, 1, 5},
// 		{1, 2, 3},
// 		{2, 3, 2},
// 		{0, 3, 10},
// 	}

// 	matrix := BuildWeightedMatrix(size, graph)

// 	// Print adjacency matrix

// 	fmt.Println("Adjacency Matrix:")
// 	for _, row := range matrix {
// 		fmt.Println(row)
// 	}
// }

//-----------------------------------------------------------------------------------------------------------------------------

func PrintMatrix(matrix [][]int) {

	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val," ")
		}
		fmt.Println()
	}
}


//-----------------------------------------------------------------------------------------------------------------------------


type Graph struct {
	size   int
	matrix [][]int
}

// Create a new graph with n vertices
func NewGraph(n int) *Graph {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return &Graph{size: n, matrix: matrix}
}



// Add edge (directed)
func (g *Graph) AddEdge(u, v int) {
	if u < 0 || v < 0 || u >= g.size || v >= g.size {
		fmt.Println("Invalid edge")
		return
	}
	g.matrix[u][v] = 1
}



// Add edge (undirected)
func (g *Graph) AddUndirectedEdge(u, v int) {
	g.AddEdge(u, v)
	g.AddEdge(v, u)
}



// Remove edge
func (g *Graph) RemoveEdge(u, v int) {
	if u < 0 || v < 0 || u >= g.size || v >= g.size {
		fmt.Println("Invalid edge")
		return
	}
	g.matrix[u][v] = 0
}



// Check if edge exists
func (g *Graph) HasEdge(u, v int) bool {
	return g.matrix[u][v] == 1
}



// Print matrix
func (g *Graph) Print() {
	for i := 0; i < g.size; i++ {
		fmt.Println(g.matrix[i])
	}
}



// func main() {
// 	g := NewGraph(3)

// 	// Add edges
// 	g.AddUndirectedEdge(0, 1)
// 	g.AddEdge(1, 2)

// 	// Print matrix
// 	g.Print()

// 	// Check edge
// 	fmt.Println("Edge 1->2 exists?", g.HasEdge(1, 2))
// }



