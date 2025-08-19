package main

import "fmt"

func solve(arr []int) []int {
	n := len(arr)
	stack := []int{}
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = arr[stack[len(stack)-1]]
		}

		stack = append(stack, i)
	}

	return res
}

func main() {
	arr := []int{4, 8, 5, 2, 25}
	fmt.Println(solve(arr))

	arr = []int{10, 9, 8, 7}
	fmt.Println(solve(arr))
}
