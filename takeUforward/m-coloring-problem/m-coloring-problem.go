package main

import "fmt"

func isValid(adj [][]int, color []int, node int, colorNumber int) bool {
	for _, v := range adj[node] {
		if color[v] == colorNumber {
			return false
		}
	}

	return true
}

func dfs(color []int, adj [][]int, N, M, node int) bool {
	if node == N {
		return true
	}

	for i := 1; i <= M; i++ {
		if isValid(adj, color, node, i) {
			color[node] = i
			if dfs(color, adj, N, M, node+1) {
				return true
			}
			color[node] = 0
		}
	}

	return false
}

func createAdjList(N int, Edges [][]int) [][]int {
	adj := make([][]int, N)
	for _, e := range Edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	return adj
}

func main() {
	N := 4
	M := 3
	Edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {0, 2}}
	adj := createAdjList(N, Edges)
	color := make([]int, N)
	fmt.Println(dfs(color, adj, N, M, 0))

	N = 3
	M = 2
	Edges = [][]int{{0, 1}, {1, 2}, {0, 2}}
	adj = createAdjList(N, Edges)
	color = make([]int, N)
	fmt.Println(dfs(color, adj, N, M, 0))
}
