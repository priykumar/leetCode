package main

import "fmt"

func insertIntoStack(stack *[]int, ele int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= ele { // descending
		*stack = append(*stack, ele)
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	insertIntoStack(stack, ele)

	*stack = append(*stack, top)
}

func sortStack(stack *[]int) {
	if len(*stack) <= 1 {
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	sortStack(stack)
	insertIntoStack(stack, top)
}

func main() {
	stack := []int{4, 1, 3, 2}
	sortStack(&stack)
	fmt.Println(stack)
}
