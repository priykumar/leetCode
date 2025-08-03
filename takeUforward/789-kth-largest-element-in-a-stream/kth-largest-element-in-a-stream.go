type MinHeap []int
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } // Min-heap: parent < children
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

type KthLargest struct {
    minheap MinHeap
    k int
}


func Constructor(k int, nums []int) KthLargest {
    minheap := &MinHeap{}
    heap.Init(minheap)

    for i:=0; i<k && i<len(nums); i++ {
        heap.Push(minheap, nums[i])
    }

    if k < len(nums) {
        for j := k; j<len(nums); j++ {
            if nums[j] > (*minheap)[0] {
                heap.Pop(minheap)
                heap.Push(minheap, nums[j])
            }
        }
    }

    fmt.Println(minheap)
    return KthLargest {
        minheap: *minheap,
        k: k,
    }
}


func (this *KthLargest) Add(val int) int {
    if this.minheap.Len() < this.k {
        heap.Push(&this.minheap, val)
    } else if val > this.minheap[0] {
        heap.Pop(&this.minheap)
        heap.Push(&this.minheap, val)
    }

    return this.minheap[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */