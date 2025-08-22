package main

import (
	"sort"
)

func minOperations(grid [][]int, x int) int {
	list := []int{}
	for _, row := range grid {
		for _, num := range row {
			list = append(list, num)
		}
	}

	sort.Ints(list)
	mid := list[len(list)/2]
	operations := 0

	for _, num := range list {
		if num < mid {
			op := mid - num
			if op%x == 0 {
				operations += op / x
			} else {
				return -1
			}
		} else if num > mid {
			op := num - mid
			if op%x == 0 {
				operations += op / x
			} else {
				return -1
			}
		}
	}

	return operations
}