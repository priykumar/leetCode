package main

import "fmt"

func dfs(u int, adj [][]int, vstd []bool, stack *[]int) {
	vstd[u] = true

	for _, v := range adj[u] {
		if !vstd[v] {
			dfs(v, adj, vstd, stack)
		}
	}

	*stack = append(*stack, u)
}

func DFS_Tsort(V int, adj [][]int) []int {
	vstd := make([]bool, V)
	stack := make([]int, 0, V)

	for u := 0; u < V; u++ {
		if !vstd[u] {
			dfs(u, adj, vstd, &stack)
		}
	}

	l, r := 0, len(stack)-1
	for l < r {
		stack[l], stack[r] = stack[r], stack[l]
		l++
		r--

	}

	return stack
}

func BFS_Tsort(V int, adj [][]int) []int {
	indegree := make([]int, V)
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(adj[i]); j++ {
			indegree[adj[i][j]]++
		}
	}

	// fmt.Println(adj, indegree)
	q := []int{}
	for i, ele := range indegree {
		if ele == 0 {
			q = append(q, i)
		}
	}

	tsort := []int{}
	for len(q) > 0 {
		u := q[0]
		// fmt.Println(u)
		q = q[1:]
		tsort = append(tsort, u)
		for _, v := range adj[u] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return tsort
}
func main() {
	V := 6
	adj := [][]int{{}, {}, {3}, {1}, {0, 1}, {0, 2}}
	tsort := BFS_Tsort(V, adj)
	fmt.Println("BFS: ", tsort)

	tsort = DFS_Tsort(V, adj)
	fmt.Println("DFS: ", tsort)
}
