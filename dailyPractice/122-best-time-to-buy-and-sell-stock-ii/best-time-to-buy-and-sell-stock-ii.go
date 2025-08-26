// func maxProfit(prices []int) int {
//     k:=1
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
//                 curr[trans] = max(-prices[i]+prev[trans-1], prev[trans])
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

    dp:=[][]int{}
    for i:=0; i<n; i++ {
        dp=append(dp, []int{-1,-1})
    }
    return solve(prices, 0, 1, dp)
}

func max(a, b int) int {
    if a>=b {
        return a
    }
    return b
}