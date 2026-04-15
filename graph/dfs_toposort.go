package graph

func topoSort(V int, adj [][]int) []int {
	visited := make([]bool, V)
	stack := []int{}

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, nei := range adj[node] {
			if !visited[nei] {
				dfs(nei)
			}
		}

		// push AFTER visiting neighbors
		stack = append(stack, node)
	}

	// run DFS for all nodes
	for i := 0; i < V; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	// reverse stack
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

// func main() {
// 	V := 6
// 	adj := [][]int{
// 		{1, 2}, // 0
// 		{3},    // 1
// 		{3},    // 2
// 		{4},    // 3
// 		{5},    // 4
// 		{},     // 5
// 	}

// 	fmt.Println(topoSort(V, adj))
// }