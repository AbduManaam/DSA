package graph

//DFS
func detectCycleDirected(n int, graph [][]int) bool {
	visited := make([]bool, n)
	path := make([]bool, n)

	var dfs func(int) bool

	dfs = func(node int) bool {
		visited[node] = true
		path[node] = true

		for _, neighbor := range graph[node] {

			// If not visited, explore
			if !visited[neighbor] {
				if dfs(neighbor) {
					return true
				}
			}

			// If already in current path → cycle
			if path[neighbor] {
				return true
			}
		}

		// Backtrack
		path[node] = false
		return false
	}

	// Check all components
	for i := 0; i < n; i++ {
		if !visited[i] {
			if dfs(i) {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	graph := [][]int{
// 		{1},     // 0 → 1
// 		{2},     // 1 → 2
// 		{3},     // 2 → 3
// 		{1},     // 3 → 1 (cycle)
// 	}

// 	fmt.Println(detectCycleDirected(4, graph)) // true
// }

