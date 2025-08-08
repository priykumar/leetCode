package main

import (
	"container/heap"
	"fmt"
)

type S struct {
	dist int
	node int
}

type MinHeap []S

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].dist < h[j].dist } // Min-heap by dist
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(S))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func prim(V int, adj [][]S) {
	parent := make([]int, V)
	key := make([]int, V)
	for i := 0; i < V; i++ {
		key[i] = 1 << 32
	}
	mst := make([]bool, V)
	for i := 0; i < V; i++ {
		parent[i] = -1
	}

	minheap := &MinHeap{}
	heap.Init(minheap)
	key[0] = 0
	heap.Push(minheap, S{0, 0})

	for i := 0; i < V-1; i++ {
		uDetail := heap.Pop(minheap).(S)
		u, _ := uDetail.node, uDetail.dist
		mst[u] = true

		for _, v := range adj[u] {
			if !mst[v.node] && key[v.node] > v.dist {
				key[v.node] = v.dist
				heap.Push(minheap, S{dist: key[v.node], node: v.node})
				parent[v.node] = u
			}
		}
	}

	fmt.Println(key)
}

func createAdjlist(V int, E [][]int) [][]S {
	adj := make([][]S, V)
	for i := 0; i < len(E); i++ {
		adj[E[i][0]] = append(adj[E[i][0]], S{node: E[i][1], dist: E[i][2]})
		adj[E[i][1]] = append(adj[E[i][1]], S{node: E[i][0], dist: E[i][2]})
	}

	return adj
}

func main() {
	V := 4
	E := [][]int{
		{0, 1, 1},
		{1, 2, 2},
		{2, 3, 3},
		{0, 3, 4},
	}
	adj := createAdjlist(V, E)
	prim(V, adj)

	V = 3
	E = [][]int{
		{0, 1, 5},
		{1, 2, 10},
		{2, 0, 15},
	}
	adj = createAdjlist(V, E)
	prim(V, adj)
}
