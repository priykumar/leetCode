// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type S struct {
// 	start  int
// 	end    int
// 	profit int
// }

// var dp []int

// func solve(idx, currEnd, n, profitTillNow int, list []S) int {
// 	if idx >= n {
// 		// fmt.Println("here", idx, profitTillNow)
// 		return profitTillNow //list[n-1].profit
// 	}

// 	if dp[idx] != -1 {
// 		return dp[idx]
// 	}

// 	// fmt.Println("start=", idx, profitTillNow)
// 	tempProfit := 0
// 	for i := idx; i < n; i++ {
// 		if list[i].start >= currEnd {
// 			tempProfit = max(tempProfit, solve(i+1, list[i].end, n, profitTillNow+list[i].profit, list))
// 		}
// 	}

// 	dp[idx] = tempProfit
// 	// fmt.Println("end=", idx, profitTillNow, tempProfit)
// 	return tempProfit
// }

package main

import (
	"fmt"
	"sort"
)

type S struct {
	start  int
	end    int
	profit int
}

var dp []int

func solve(idx int, list []S) int {
	if idx >= len(list) {
		return 0
	}
	if dp[idx] != -1 {
		return dp[idx]
	}

	// not take
	notTake := solve(idx+1, list)

	// take
	next := idx + 1
	for next < len(list) && list[next].start < list[idx].end {
		next++
	}
	take := list[idx].profit + solve(next, list)

	dp[idx] = max(take, notTake)
	return dp[idx]
}

func main() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	dp = make([]int, len(startTime))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	list := []S{}
	for i := 0; i < len(startTime); i++ {
		list = append(list, S{startTime[i], endTime[i], profit[i]})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].start < list[j].start
	})
	//fmt.Println(solve(0, 0, len(startTime), 0, list))
	fmt.Println(solve(0, list))

	startTime = []int{1, 2, 3, 3}
	endTime = []int{3, 4, 5, 6}
	profit = []int{50, 10, 40, 70}
	dp = make([]int, len(startTime))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	list = []S{}
	for i := 0; i < len(startTime); i++ {
		list = append(list, S{startTime[i], endTime[i], profit[i]})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].start < list[j].start
	})

	//fmt.Println(solve(0, 0, len(startTime), 0, list))
	fmt.Println(solve(0, list))

}
