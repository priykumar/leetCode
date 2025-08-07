package main

import "fmt"

var res []int

func dfs(node int, vstd []bool, adj [][]int) {
	vstd[node] = true

	for _, v := range adj[node] {
		if !vstd[v] {
			res = append(res, v)
			dfs(v, vstd, adj)
		}
	}
}

func main() {
	V := 5
	adj := [][]int{
		{2, 3, 1},
		{0},
		{0, 4},
		{0},
		{2},
	}

	vstd := make([]bool, V)
	for u := 0; u < V; u++ {
		if vstd[u] == false {
			res = append(res, u)
			dfs(u, vstd, adj)
			fmt.Println("DFS: ", res)
		}
	}

	res = []int{0}
	vstd = make([]bool, V)
	q := []int{0}
	vstd[0] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		for _, v := range adj[u] {
			if !vstd[v] {
				vstd[v] = true
				q = append(q, v)
				res = append(res, v)
			}
		}
	}

	fmt.Println("BFS: ", res)

}
