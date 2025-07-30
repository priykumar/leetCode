package main

var dp [][]int

func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := weights[0]; i <= capacity; i++ {
		dp[0][i] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= capacity; j++ {
			// not take
			nottake_val := dp[i-1][j]
			//take
			take_val := 0
			if weights[i] <= j {
				take_val = dp[i-1][j-weights[i]] + values[i]
			}

			dp[i][j] = max(nottake_val, take_val)
		}
	}
	return dp[n-1][capacity]
}

func knapsack_space_optimised(weights []int, values []int, capacity int) int {
	n := len(weights)

	curr := make([]int, capacity+1)
	prev := make([]int, capacity+1)

	for i := weights[0]; i <= capacity; i++ {
		prev[i] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= capacity; j++ {
			// not take
			nottake_val := prev[j]
			//take
			take_val := 0
			if weights[i] <= j {
				take_val = prev[j-weights[i]] + values[i]
			}

			curr[j] = max(nottake_val, take_val)
		}
		copy(prev, curr)
	}
	return prev[capacity]
}

func main() {
	weights := []int{10, 20, 30}
	values := []int{60, 100, 120}
	capacity := 50
	result := knapsack(weights, values, capacity)
	println("Maximum value in Knapsack =", result)

	weights = []int{5, 4, 6, 3}
	values = []int{10, 40, 30, 50}
	capacity = 10
	result = knapsack(weights, values, capacity)
	println("Maximum value in Knapsack =", result)

	weights = []int{1, 2, 3, 8, 7, 4}
	values = []int{20, 5, 10, 40, 15, 25}
	capacity = 10
	result = knapsack(weights, values, capacity)
	println("Maximum value in Knapsack =", result)
}
