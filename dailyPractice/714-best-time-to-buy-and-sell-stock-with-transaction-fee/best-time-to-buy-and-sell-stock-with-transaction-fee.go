func solve(prices []int, fee int, isBuy int, i int, dp [][]int) int {
    if i >= len(prices) {
        return 0
    }

    if dp[i][isBuy] != -1 {
        return dp[i][isBuy]
    }

    if isBuy == 1 {
        dp[i][isBuy] = max(-prices[i]+solve(prices, fee, 0, i+1, dp), solve(prices, fee, 1, i+1, dp))
    } else {
        dp[i][isBuy] = max(-fee+prices[i]+solve(prices, fee, 1, i+1, dp), solve(prices, fee, 0, i+1, dp))
    }

    return dp[i][isBuy]
}

func maxProfit(prices []int, fee int) int {
    n := len(prices)
    dp := make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i]=[]int{-1, -1}
    }

    return solve(prices, fee, 1, 0, dp)
}