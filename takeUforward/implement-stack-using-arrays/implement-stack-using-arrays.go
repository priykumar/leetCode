package main

import "fmt"

type stack struct {
	arr    []int
	maxlen int
}

func ArrayStack() stack {
	return stack{
		arr:    make([]int, 0),
		maxlen: 5,
	}
}

func (s *stack) push(ele int) {
	if s.len() < s.maxlen {
		s.arr = append(s.arr, ele)
	}
}

func (s *stack) pop() int {
	ele := -1
	if !s.isEmpty() {
		n := s.len()
		ele = s.arr[n-1]
		s.arr = s.arr[:n-1]
	}
	return ele
}

func (s stack) top() int {
	ele := -1
	if !s.isEmpty() {
		n := s.len()
		ele = s.arr[n-1]
	}
	return ele
}

func (s stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s stack) len() int {
	return len(s.arr)
}

func main() {
	s := ArrayStack()
	fmt.Println(s.isEmpty())
	s.push(1)
	fmt.Println(s.pop())
	fmt.Println(s.len())

	s.push(1)
	s.push(1)
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(7)

	fmt.Println(s.len())
	fmt.Println(s.top())

	fmt.Println(s.pop())
	fmt.Println(s.top())
}
