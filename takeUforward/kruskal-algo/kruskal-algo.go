package main

import (
	"fmt"
	"sort"
)

var parent []int
var rank []int

func findParent(u int) int {
	if u == parent[u] {
		return u
	}
	parent[u] = findParent(parent[u])
	return parent[u]
}

func union(u, v int) {
	pu, pv := findParent(u), findParent(v)
	if pu == pv {
		return
	} else if rank[pu] > rank[pv] {
		parent[pv] = pu
	} else if rank[pu] < rank[pv] {
		parent[pu] = pv
	} else {
		parent[pv] = pu
		rank[pu]++
	}
}

func main() {
	// weight, u, v
	adj := [][]int{
		{9, 4, 5},
		{4, 1, 5},
		{1, 1, 4},
		{5, 4, 3},
		{3, 2, 4},
		{2, 1, 2},
		{3, 3, 2},
		{8, 3, 6},
		{7, 2, 6},
	}

	rank = make([]int, 7)
	parent = make([]int, 7)
	for i := 1; i < 7; i++ {
		parent[i] = i
	}

	sort.SliceStable(adj, func(i, j int) bool {
		return adj[i][0] < adj[j][0]
	})

	mst := 0
	for i := 0; i < len(adj); i++ {
		w, u, v := adj[i][0], adj[i][1], adj[i][2]
		if findParent(u) != findParent(v) {
			mst += w
			union(u, v)
		}
	}

	fmt.Println(mst)
	fmt.Println(parent)
}
