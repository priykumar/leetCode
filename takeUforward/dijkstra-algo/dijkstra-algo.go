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

const INT_MAX int = 1 << 32

func dijkstra(src, V int, adj [][][]int) []int {
	minheap := &MinHeap{}
	heap.Init(minheap)
	heap.Push(minheap, S{src, 0})

	minDist := make([]int, V)
	for i := 0; i < V; i++ {
		minDist[i] = INT_MAX
	}
	minDist[src] = 0

	for minheap.Len() > 0 {
		uDetail := heap.Pop(minheap).(S)
		du, u := uDetail.dist, uDetail.node

		for _, v := range adj[u] {
			if minDist[v[0]] > du+v[1] {
				minDist[v[0]] = du + v[1]
				heap.Push(minheap, S{du + v[1], v[0]})
			}
		}
	}

	return minDist
}

func main() {
	V := 2
	adj := [][][]int{
		{{1, 9}}, // node, dist
		{{0, 9}},
	}
	src := 0
	fmt.Println(dijkstra(src, V, adj))

	V = 3
	adj = [][][]int{
		{{1, 1}, {2, 6}}, // node, dist
		{{2, 3}, {0, 1}},
		{{1, 3}, {0, 6}},
	}
	src = 2
	fmt.Println(dijkstra(src, V, adj))

	V = 8
	adj = [][][]int{
		{{1, 7}, {2, 8}, {7, 10}}, // node, dist
		{{0, 7}, {3, 6}, {2, 3}},
		{{1, 3}, {0, 8}, {7, 1}, {4, 3}, {3, 4}},
		{{1, 6}, {2, 4}, {4, 2}, {5, 5}},
		{{3, 2}, {5, 5}, {6, 2}, {7, 6}, {2, 3}},
		{{3, 5}, {4, 5}, {6, 3}},
		{{5, 3}, {4, 2}, {7, 1}},
		{{0, 10}, {2, 1}, {4, 6}, {6, 1}},
	}
	src = 0
	fmt.Println(dijkstra(src, V, adj))
}
