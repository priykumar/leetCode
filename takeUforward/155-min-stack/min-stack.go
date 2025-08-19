type MinStack struct {
    arr []int
    minArrIdx []int
}


func Constructor() MinStack {
    return MinStack {
        arr: make([]int, 0),
        minArrIdx: make([]int, 0),
    }
}


func (this *MinStack) Push(val int)  {
    this.arr = append(this.arr, val)

    // Insert idx of ele from arr to minArr
    if len(this.minArrIdx) == 0 || val < this.arr[this.minArrIdx[len(this.minArrIdx)-1]] {
        this.minArrIdx = append(this.minArrIdx, len(this.arr)-1)
    }
}


func (this *MinStack) Pop()  {
    // ele := this.arr[len(this.arr)-1]
    // Remove entry from minArr
    IdxOfEleToBePopped := len(this.arr)-1
    if this.minArrIdx[len(this.minArrIdx)-1] == IdxOfEleToBePopped {
        this.minArrIdx = this.minArrIdx[:len(this.minArrIdx)-1]
    }

    this.arr = this.arr[:len(this.arr)-1]
}


func (this *MinStack) Top() int {
    return this.arr[len(this.arr)-1]
}


func (this *MinStack) GetMin() int {
    return this.arr[this.minArrIdx[len(this.minArrIdx)-1]]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */