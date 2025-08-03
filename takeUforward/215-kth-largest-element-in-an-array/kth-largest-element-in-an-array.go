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

func findKthLargest(nums []int, k int) int {
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

    return (*minheap)[0] 
}