package main

import (
	"fmt"
	"sort"
)

func solve(Jobs [][]int) {
	sort.SliceStable(Jobs, func(i, j int) bool {
		return Jobs[i][2] > Jobs[j][2]
	})

	maxDeadline := Jobs[0][1]
	for i := 1; i < len(Jobs); i++ {
		maxDeadline = max(maxDeadline, Jobs[i][1])
	}

	arr := make([]int, maxDeadline+1)
	doneCount, maxProfit := 0, 0
	for _, j := range Jobs {
		if arr[j[1]] == 0 {
			arr[j[1]] = j[0]
			doneCount++
			maxProfit += j[2]
		} else {
			i := j[1]
			for i > 0 {
				if arr[i] == 0 {
					arr[i] = j[0]
					doneCount++
					maxProfit += j[2]
					break
				}
				i--
			}
		}
	}

	fmt.Println(doneCount, maxProfit)
}

func main() {
	Jobs := [][]int{{1, 4, 20}, {2, 1, 10}, {3, 1, 40}, {4, 1, 30}}
	solve(Jobs)

	Jobs = [][]int{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 1, 15}}
	solve(Jobs)
}
