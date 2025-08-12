const INT_MAX int = 1<<32
func recursiveApproach(l, r int, cuts []int, dp [][]int) int{
    if l > r {
        return 0
    }
    if dp[l][r] != -1 {
        return dp[l][r]
    }

    res := INT_MAX
    for idx:=l; idx<=r; idx++ {
        tempRes := cuts[r+1]-cuts[l-1] + 
        recursiveApproach(l, idx-1, cuts, dp) + 
        recursiveApproach(idx+1, r, cuts, dp)
        res = min(res, tempRes)
    }

    dp[l][r] = res
    return res
}

func minCost(n int, cuts []int) int {
    sz := len(cuts)
    cuts = append([]int{0}, cuts...)
    cuts = append(cuts, n)
    sort.Ints(cuts)

    dp := make([][]int, sz+1)
    for i:=0; i<sz+1; i++ {
        dp[i] = make([]int, sz+1)
        for j:=0; j<sz+1; j++ {
            dp[i][j]=-1
        }
    }
    res := recursiveApproach(1, sz, cuts, dp)
    return res
}