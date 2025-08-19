package main

import "fmt"

func solve(arr []int) {
	n := len(arr)

	// Find first min at left
	stack := []int{}
	left := make([]int, n)
	left[0] = -1

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Find first min at right
	stack = []int{}

	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		window := right[i] - left[i] - 1
		res[window-1] = max(res[window-1], arr[i])
	}

	// If a value X is the answer for some larger window, then for all smaller windows the answer must be at least X.

	for i := n - 2; i >= 0; i-- {
		res[i] = max(res[i], res[i+1])
	}
	fmt.Println(res)
}

func main() {
	arr := []int{10, 20, 30, 50, 10, 70, 30}
	solve(arr)

	arr = []int{6, 3, 5, 1, 12}
	solve(arr)
}

// This problem is for max of mins for every window
// But if question would have been min of mins -> Same answer for every window = global min of the array
// But if question would have been min of maxs -> Nearest greater on left & right (stack) → reverse pass with min
// But if question would have been max of maxs -> Same answer for every window = global max of the array
