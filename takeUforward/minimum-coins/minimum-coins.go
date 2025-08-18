package main

import (
	"fmt"
	"sort"
)

func solve(coins []int, amount int) int {
	n := len(coins)
	sort.Ints(coins)
	dp := make([][]int, n)
	temparr := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		temparr[i] = 1 << 32
	}

	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
		copy(dp[i], temparr)
	}

	for i := coins[0]; i <= amount; i = i + coins[0] {
		dp[0][i] = i / coins[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= amount; j++ {
			// nottake
			dp[i][j] = dp[i-1][j]
			// take
			if j >= coins[i] {
				dp[i][j] = min(dp[i][j], dp[i][j-coins[i]]+1)
			}
		}
	}

	if dp[n-1][amount] == 1<<32 {
		return -1
	}
	return dp[n-1][amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(solve(coins, amount))

	coins = []int{2, 5}
	amount = 3
	fmt.Println(solve(coins, amount))

	coins = []int{1}
	amount = 0
	fmt.Println(solve(coins, amount))
}
