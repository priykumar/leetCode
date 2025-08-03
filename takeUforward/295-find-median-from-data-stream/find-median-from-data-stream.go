type MinHeap []int
type MaxHeap []int

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

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } // Max-heap: parent > children
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[0 : n-1]
	return val
}

type MedianFinder struct {
    minheap MinHeap
    maxheap MaxHeap
}


func Constructor() MedianFinder {
    minheap := &MinHeap{}
	heap.Init(minheap)

    maxheap := &MaxHeap{}
	heap.Init(maxheap)

    return MedianFinder {
        minheap: *minheap,
        maxheap: *maxheap,
    }
}


func (this *MedianFinder) AddNum(num int)  {
    if this.maxheap.Len() == 0 || num < this.maxheap[0] {
        heap.Push(&this.maxheap, num)
        if this.maxheap.Len() - this.minheap.Len() > 1 {
            ele := heap.Pop(&this.maxheap)
            heap.Push(&this.minheap, ele)
        }
    } else {
        heap.Push(&this.minheap, num)
        if this.maxheap.Len() < this.minheap.Len() {
            ele := heap.Pop(&this.minheap)
            heap.Push(&this.maxheap, ele)
        }
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if (this.maxheap.Len() + this.minheap.Len()) % 2 == 0 {
        return float64(this.maxheap[0] + this.minheap[0])/2
    }

    return float64(this.maxheap[0])
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */