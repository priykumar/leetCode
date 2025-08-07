package main

import "fmt"

func bfs_check_cycle(u int, adj [][]int, vstd []bool) bool {
	type S struct {
		node   int
		parent int
	}

	vstd[u] = true
	q := []S{{u, -1}} // node, parent node

	for len(q) > 0 {
		u = q[0].node
		parent := q[0].parent
		q = q[1:]
		for _, v := range adj[u] {
			if !vstd[v] {
				q = append(q, S{v, u})
				vstd[v] = true
			} else if v != parent {
				return true
			}
		}
	}

	return false
}

func BFS_Approach(V int, adj [][]int) bool {
	vstd := make([]bool, V)
	for i := 0; i < V; i++ {
		if !vstd[i] {
			if bfs_check_cycle(i, adj, vstd) {
				return true
			}
		}
	}

	return false
}

func dfs_check_cycle(u, parent int, adj [][]int, vstd []bool) bool {
	vstd[u] = true

	for _, v := range adj[u] {
		if !vstd[v] {
			if dfs_check_cycle(v, u, adj, vstd) {
				return true
			}
		} else if v != parent {
			return true
		}
	}

	return false
}

func DFS_Approach(V int, adj [][]int) bool {
	vstd := make([]bool, V)
	for u := 0; u < V; u++ {
		if !vstd[u] {
			if dfs_check_cycle(u, -1, adj, vstd) {
				return true
			}
		}
	}

	return false
}

func main() {
	V := 6
	adj := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	fmt.Println("BFS: ", BFS_Approach(V, adj))
	fmt.Println("DFS: ", DFS_Approach(V, adj))

	V = 4
	adj = [][]int{
		{1, 2},
		{0},
		{0, 3},
		{2},
	}
	fmt.Println("BFS: ", BFS_Approach(V, adj))
	fmt.Println("DFS: ", DFS_Approach(V, adj))

}
