package main

import (
	"fmt"
	"sort"
)

// maximum possible minimum distance -> Need to find the min distance

func countCowsFor_m_distance(stalls []int, m, n int) int {
	cow := 1
	lastUseStall := 0
	for i := 1; i < n; i++ {
		if stalls[i]-stalls[lastUseStall] >= m {
			cow++
			lastUseStall = i
		}
	}

	return cow
}

func solve(n, cows int, stalls []int) int {
	sort.Ints(stalls)
	l, r := 1, stalls[n-1]-stalls[0]

	// reqCow > cows -> l = m+1
	// reqCow < cows -> r=m-1
	// reqCow = cows -> l =m+1
	res := 0
	for l <= r {
		m := l + (r-l)/2
		reqCow := countCowsFor_m_distance(stalls, m, n)
		if reqCow >= cows {
			l = m + 1
			res = m
		} else {
			r = m - 1
		}
	}

	return res
}

func main() {
	n := 6
	cows := 4
	stalls := []int{0, 3, 4, 7, 10, 9}
	fmt.Println(solve(n, cows, stalls))

	n = 5
	cows = 2
	stalls = []int{4, 2, 1, 3, 6}
	fmt.Println(solve(n, cows, stalls))

	n = 5
	cows = 3
	stalls = []int{10, 1, 2, 7, 5}
	fmt.Println(solve(n, cows, stalls))
}
