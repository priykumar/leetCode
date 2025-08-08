package main

import "fmt"

func dfs(u int, vstd []bool, adj [][]int, stack *[]int) {
	vstd[u] = true
	for _, v := range adj[u] {
		if !vstd[v] {
			dfs(v, vstd, adj, stack)
		}
	}
	if stack != nil {
		*stack = append(*stack, u)
	}
}

func kosaraju(V int, adj [][]int) int {
	stack := []int{}
	vstd := make([]bool, V)

	// Tsort
	for i := 0; i < V; i++ {
		if !vstd[i] {
			dfs(i, vstd, adj, &stack)
		}
	}

	// Reverse Adj List
	revAdj := make([][]int, V)
	for i := 0; i < V; i++ {
		for _, v := range adj[i] {
			revAdj[v] = append(revAdj[v], i)
		}
	}

	vstd = make([]bool, V)
	// Perform DFS
	count := 0
	for i := len(stack) - 1; i >= 0; i-- {
		if vstd[stack[i]] == false {
			dfs(stack[i], vstd, revAdj, nil)
			count++
		}
	}

	return count
}

func createAdjList(V int, E [][]int) [][]int {
	adj := make([][]int, V)
	for i := 0; i < len(E); i++ {
		u, v := E[i][0], E[i][1]
		adj[u] = append(adj[u], v)
	}

	return adj
}

func main() {
	V := 5
	E := [][]int{
		{0, 2},
		{1, 0},
		{2, 1},
		{0, 3},
		{3, 4},
	}
	adj := createAdjList(V, E)
	fmt.Println(kosaraju(V, adj))

	V = 8
	E = [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{2, 3},
		{3, 4},
		{4, 5},
		{4, 7},
		{5, 6},
		{6, 4},
		{6, 7},
	}
	adj = createAdjList(V, E)
	fmt.Println(kosaraju(V, adj))
}
