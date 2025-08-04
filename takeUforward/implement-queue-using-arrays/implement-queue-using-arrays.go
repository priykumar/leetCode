package main

import "fmt"

type queue struct {
	arr    []int
	front  int // delete from front
	rear   int // insert in rear
	maxlen int
}

func Queue() queue {
	return queue{
		arr:    make([]int, 0),
		front:  -1,
		rear:   0,
		maxlen: 5,
	}
}

func (q *queue) push(ele int) {
	if q.len() < q.maxlen {
		q.arr = append(q.arr, ele)
	}
}

func (q *queue) pop() int {
	ele := -1
	if !q.isEmpty() {
		ele = q.arr[0]
		q.arr = q.arr[1:]
	}
	return ele
}

func (q queue) peek() int {
	ele := -1
	n := q.len()
	if n > 0 {
		ele = q.arr[n-1]
	}
	return ele
}

func (q queue) isEmpty() bool {
	return q.len() == 0
}

func (q queue) len() int {
	return len(q.arr)
}

func main() {
	q := Queue()
	fmt.Println(q.isEmpty())
	q.push(1)
	fmt.Println(q.pop())

	q.push(1)
	q.push(10)
	q.push(2)
	q.push(3)
	q.push(4)
	q.push(5)
	q.push(6)

	fmt.Println(q.arr)
	fmt.Println(q.peek())
	fmt.Println(q.pop())
	fmt.Println(q.peek())
	fmt.Println(q.arr)
}
