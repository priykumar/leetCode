func isInterleave(s1 string , s2 string , s3 string )  (bool) {
    a, b, c := len(s1), len(s2), len(s3)
    if a+b!=c {
        return false
    }
    dp:=make([][]bool, a+1)
    for i:=0; i<=a; i++ {
        dp[i]=make([]bool, b+1)
    }

    dp[0][0]=true
    for i:=1; i<=a; i++ {
        if s3[i-1]==s1[i-1] {
            dp[i][0]=dp[i-1][0]
        }
    }
    for i:=1; i<=b; i++ {
        if s3[i-1]==s2[i-1] {
            dp[0][i]=dp[0][i-1]
        }
    }

    for i:=1; i<=a; i++ {
        for j:=1; j<=b; j++ {
            a, b :=false, false
            if s3[i+j-1]==s1[i-1] {
                a = dp[i-1][j]
            } 
            if s3[i+j-1]==s2[j-1] {
                b = dp[i][j-1]
            }
            dp[i][j] = a || b
        }
    }

    return dp[a][b]
}
