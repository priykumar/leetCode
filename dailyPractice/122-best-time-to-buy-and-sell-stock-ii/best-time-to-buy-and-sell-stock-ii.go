// func maxProfit(prices []int) int {
//     prev:=[]int{-1*prices[0],0}
//     curr:=[]int{0,0}
//     for i:=1; i<len(prices); i++ {
//         for sell:=0; sell<2; sell++ {
//             if sell==0 {
//                 // buy
//                 curr[sell]=max(-1*prices[i]+prev[1], prev[0])
//             } else {
//                 // sell
//                 curr[sell]=max(prices[i]+prev[0], prev[1])
//             }
//         }
//         copy(prev, curr)
//         //fmt.Println(i, prev[0], prev[1])
//     }
//     return prev[1]
// }

func solve(a[]int, i int, isBuy int, dp [][]int) int {
    if i>=len(a) {
        return 0
    }

    if (dp[i][isBuy]!=-1) {
        return dp[i][isBuy]
    }
    if isBuy==1 {
        dp[i][isBuy]=max(-a[i]+solve(a, i+1, 0, dp), solve(a, i+1, 1, dp))
    } else {
        dp[i][isBuy]=max(a[i]+solve(a, i+1, 1, dp), solve(a, i+1, 0, dp))
    }

    return dp[i][isBuy]
}

func maxProfit(prices []int) int {
    n:=len(prices)
    if n<=1 {
        return 0
    }

    dp:=make([][]int, n)
    for i:=0; i<n; i++ {
        dp[i]=[]int{-1,-1}
    }
    return solve(prices, 0, 1, dp)
}

// func max(a, b int) int {
//     if a>=b {
//         return a
//     }
//     return b
// }