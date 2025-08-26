// func maxProfit(k int, prices []int) int {
//     prev, curr := make([]int, 2*k), make([]int, 2*k)
//     for t := 0; t < 2*k; t++ {
// 		if t%2 == 0 { // buy state
// 			prev[t] = -prices[0]
// 		}
// 	}

//     for i := 1; i < len(prices); i++ {
//         for trans:=0; trans<2*k; trans++ {
//             if trans%2 == 0{
//                 // buy
//                 if trans ==0 {
//                     curr[trans] = max(-prices[i], prev[trans])
//                 } else {
//                     curr[trans] = max(-prices[i]+prev[trans-1], prev[trans])
//                 }
//             } else {
//                 curr[trans] = max(prices[i]+prev[trans-1], prev[trans])
//             }
//         }
//         copy(prev, curr)
//     }

//     // check all sell states
//     res := 0
//     for i:=1; i<2*k; i+=2 {
//         res = max(res, prev[i])
//     }
//     return res
// }

func solve(a []int, i int, dp [][]int, trans, k int) int {
    if i>=len(a) || trans == 2*k {
        return 0
    }

    if (dp[i][trans]!=-1) {
        return dp[i][trans]
    }
    if trans%2==0 {
        dp[i][trans]=max(-a[i]+solve(a, i+1, dp, trans+1, k), solve(a, i+1, dp, trans, k))
    } else {
        dp[i][trans]=max(a[i]+solve(a, i+1, dp, trans+1, k), solve(a, i+1, dp, trans, k))
    }

    return dp[i][trans] // the maximum profit you can achieve starting from day i, given that you have already completed trans transactions
}

func maxProfit(k int, prices []int) int {
    n:=len(prices)
    if n<=1 {
        return 0
    }

    dp:=make([][]int, n)
    temp:=make([]int, 2*k)
    for i:=0; i<2*k; i++ {
        temp[i]=-1
    }

    for i:=0; i<n; i++ {
        dp[i]=make([]int, 2*k)
        copy(dp[i], temp)
    }
    return solve(prices, 0, dp, 0, k)
}

func max(a, b int) int {
    if a>=b {
        return a
    }
    return b
}