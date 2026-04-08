package graph

import "fmt"

func BuildUndirectedMatrix(s int, graph [][]int) [][]int {

	matrix := make([][]int, s)

	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)
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

func BuildWeightedMatrix(s int, graph [][]int) [][]int {

	matrix := make([][]int, s)

	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)
	}

	for _, e := range graph {
		u, v, w := e[0], e[1], e[3]
		matrix[u][v] = w
		matrix[v][u] = w
	}
	return matrix

}

func PrintMatrix(matrix [][]int) {

	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val," ")
		}
		fmt.Println()
	}
}




