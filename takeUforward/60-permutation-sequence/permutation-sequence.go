
func solve(n, k int, currPos *int, vstd []bool, curr string, res *string) { 
    if len(curr) == n {
        *currPos=*currPos+1
        if *currPos == k {
            *res = curr
            return
        }
    }

    if *currPos > k {
        return
    }
    for i:=1; i<=n; i++ {
        if !vstd[i] {
            curr+=fmt.Sprintf("%d", i)
            vstd[i] = true
            solve(n, k, currPos, vstd, curr, res)
            vstd[i] = false
            curr=curr[:len(curr)-1]
        }
    }
}

func getPermutation(n int, k int) string {
    vstd := make([]bool, n+1)
    res := ""
    currPos := 0
    solve(n, k, &currPos, vstd, "", &res)
    return res
}