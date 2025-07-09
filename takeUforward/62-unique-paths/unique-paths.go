func uniquePaths(m int, n int) int {
    dp := [][]int{}
    temp := []int{}

    for i:=0; i<n; i++ {
        temp=append(temp, 0)
    }

    for i:=0; i<m; i++ {
        dp=append(dp, temp)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if i==0 || j==0 {
                dp[i][j]=1
                continue
            }
            dp[i][j]=dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}