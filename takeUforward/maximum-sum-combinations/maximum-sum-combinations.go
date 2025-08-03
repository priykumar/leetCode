package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type S struct {
	sum int
	x   int
	y   int
}

type MaxHeap []S

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool  { return h[i].sum > h[j].sum }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(S)) }
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	last := (*h)[n-1]
	*h = (*h)[:n-1]
	return last
}

func solve(a, b []int, k int) []int {
	sort.Ints(a)
	sort.Ints(b)

	visited := map[[2]int]bool{}
	maxheap := &MaxHeap{}
	heap.Init(maxheap)

	alen, blen := len(a), len(b)
	heap.Push(maxheap, S{a[alen-1] + b[blen-1], alen - 1, blen - 1})
	visited[[2]int{alen - 1, blen - 1}] = true

	res := []int{}
	for k > 0 {
		tempS := heap.Pop(maxheap).(S)
		res = append(res, tempS.sum)

		if visited[[2]int{tempS.x, tempS.y - 1}] == false && tempS.y-1 >= 0 {
			heap.Push(maxheap, S{a[tempS.x] + b[tempS.y-1], tempS.x, tempS.y - 1})
			visited[[2]int{tempS.x, tempS.y - 1}] = true
		}

		if visited[[2]int{tempS.x - 1, tempS.y}] == false && tempS.x-1 >= 0 {
			heap.Push(maxheap, S{a[tempS.x-1] + b[tempS.y], tempS.x - 1, tempS.y})
			visited[[2]int{tempS.x - 1, tempS.y}] = true
		}

		k--
	}

	return res
}

func main() {
	nums1 := []int{3, 4, 5}
	nums2 := []int{2, 6, 3}
	k := 2
	fmt.Println(solve(nums1, nums2, k))

	nums1 = []int{7, 3}
	nums2 = []int{1, 6}
	k = 2
	fmt.Println(solve(nums1, nums2, k))
}
