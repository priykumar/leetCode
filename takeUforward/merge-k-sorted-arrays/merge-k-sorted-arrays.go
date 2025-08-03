package main

import (
	"container/heap"
	"fmt"
)

// Element holds the value and its position in original arrays
type Element struct {
	val     int
	arrIdx  int // index of the array
	elemIdx int // index within the array
}

type MinHeap []Element

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val } // Min-Heap
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Element)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func solve(list [][]int) []int {
	// 	1. Use a min-heap to always keep track of the smallest current element across the K arrays.
	// 	2. Push the first element of each array into the heap along with:
	// 		a. The array index (to identify which array it came from)
	// 		b. The element index (to fetch the next element from the same array)
	// 	3. Pop the smallest element from the heap and push the next element from the same array into the heap.
	// 	4. Repeat until the heap is empty.

	minheap := &MinHeap{}
	heap.Init(minheap)

	// Push the first element of each array into the heap
	for i := 0; i < len(list); i++ {
		heap.Push(minheap, Element{list[i][0], i, 0})
	}

	res := []int{}
	for minheap.Len() > 0 {
		curr := heap.Pop(minheap).(Element)
		res = append(res, curr.val)

		listIdx := curr.arrIdx
		eleIdx := curr.elemIdx
		if eleIdx+1 < len(list[listIdx]) {
			heap.Push(minheap, Element{list[listIdx][eleIdx+1], listIdx, eleIdx + 1})
		}
	}

	return res
}

func main() {
	list := [][]int{
		{1, 2, 3, 4},
		{2, 2, 3, 4},
		{5, 5, 6, 6},
		{7, 8, 9, 9},
	}

	fmt.Println(solve(list))

	list = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(solve(list))
}
