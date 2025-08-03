type Item struct {
	val   int
	count int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].count < h[j].count } // Min-heap by count
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
    minheap := &MinHeap{}
    heap.Init(minheap)

    eleCount := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        eleCount[nums[i]]++
    }

    for key, val := range eleCount {
        if minheap.Len() == k {
            if (*minheap)[0].count < val {
                heap.Pop(minheap)
                heap.Push(minheap, Item{key, val})
            }
        } else {
            heap.Push(minheap, Item{key, val})
        }
    }

    res := make([]int, 0, k)
    for minheap.Len() > 0 {
        ele := heap.Pop(minheap).(Item)
        res=append(res, ele.val)
    }

    return res
}