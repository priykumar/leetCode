package main

type MaxHeap struct {
	arr []int
}

// Initialize the heap
func initializeHeap() MaxHeap {
	return MaxHeap{
		arr: make([]int, 0),
	}
}

func (this *MaxHeap) heapify(idx int) {
	largest := idx
	lc, rc := 2*idx+1, 2*idx+2
	n := this.heapSize()

	if lc < n && this.arr[largest] < this.arr[lc] {
		largest = lc
	}
	if rc < n && this.arr[largest] < this.arr[rc] {
		largest = rc
	}

	if largest != idx {
		this.arr[largest], this.arr[idx] = this.arr[idx], this.arr[largest]
		this.heapify(largest)
	}
}

// insert value x to the max heap
func (this *MaxHeap) push(key int) {
	this.arr = append(this.arr, key)

	n := this.heapSize()
	if n == 1 {
		return
	}

	i := n - 1
	for i > 0 && this.arr[(i-1)/2] < this.arr[i] {
		this.arr[(i-1)/2], this.arr[i] = this.arr[i], this.arr[(i-1)/2]
		i = (i - 1) / 2
	}
}

// update the value at given index to val (index will be given 0-based indexing)
func (this *MaxHeap) changeKey(index, new_val int) {
	old := this.arr[index]
	this.arr[index] = new_val
	if new_val > old {
		// Bubble up
		i := index
		for i > 0 && this.arr[(i-1)/2] < this.arr[i] {
			this.arr[(i-1)/2], this.arr[i] = this.arr[i], this.arr[(i-1)/2]
			i = (i - 1) / 2
		}
	} else {
		// Heapify down
		this.heapify(index)
	}
}

// Remove the maximum element from the heap
func (this *MaxHeap) pop() int {
	n := this.heapSize()
	if n == 0 {
		return -1
	} else if n == 1 {
		ele := this.arr[0]
		this.arr = this.arr[1:]
		return ele
	}

	ele := this.arr[0]
	this.arr[0] = this.arr[n-1]
	this.arr = this.arr[:n-1]
	this.heapify(0)

	return ele
}

// returns if heap is empty or not
func (this MaxHeap) isEmpty() bool {
	return len(this.arr) == 0
}

// Output the maximum value from the max heap
func (this MaxHeap) getMax() int {
	if len(this.arr) > 0 {
		return this.arr[0]
	}
	return -1
}

// return the current size of the heap
func (this MaxHeap) heapSize() int {
	return len(this.arr)
}
