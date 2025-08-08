package main

import "fmt"

const INT_MAX int = 1 << 32

func bellman(V int, E [][]int, src int) []int {
	minDist := make([]int, V)
	for i := 0; i < V; i++ {
		minDist[i] = INT_MAX
	}
	minDist[src] = 0

	for i := 0; i < V; i++ {
		for j := 0; j < len(E); j++ {
			u, v, w := E[j][0], E[j][1], E[j][2]
			if minDist[v] > minDist[u]+w {
				minDist[v] = minDist[u] + w
			}
		}
	}

	return minDist
}

func main() {
	V := 6
	Edges := [][]int{
		{3, 2, 6},
		{5, 3, 1},
		{0, 1, 5},
		{1, 5, -3},
		{1, 2, -2},
		{3, 4, -2},
		{2, 4, 3},
	}
	src := 0
	fmt.Println(bellman(V, Edges, src))

	V = 2
	Edges = [][]int{
		{0, 1, 9},
	}
	src = 0
	fmt.Println(bellman(V, Edges, src))
}
