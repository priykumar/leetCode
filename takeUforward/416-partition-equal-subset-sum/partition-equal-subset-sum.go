func tabulation(nums []int, totalSum int) []bool {
    n := len(nums)

    dp := [][]bool{}
    temp := []bool{}
    for i:=0; i<=totalSum; i++ {
        temp=append(temp, false)
    }
    for i:=0; i<n; i++ {
        temp1:=make([]bool, len(temp))
        copy(temp1, temp)
        dp=append(dp, temp1)
    }

    sort.Ints(nums)
    for i:=0; i<n; i++ {
        dp[i][nums[i]] = true
    }

    for i:=1; i<n; i++ {
        for j:=1; j<=totalSum; j++ {
            // not take
            dp[i][j] = dp[i][j] || dp[i-1][j]
            // take
            if j>=nums[i] {
                dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
            }
        }
    }

    // fmt.Println(dp)
    return dp[n-1]
}

func tabulation1(nums []int, totalSum int) []bool {
    n := len(nums)

    prev, curr := []bool{}, []bool{}
    for i:=0; i<=totalSum; i++ {
        prev=append(prev, false)
        curr=append(curr, false)
    }

    sort.Ints(nums)
    prev[nums[0]] = true

    for i:=1; i<n; i++ {
        for j:=1; j<=totalSum; j++ {
            // not take
            curr[j] = curr[j] || prev[j]
            // take
            if j>=nums[i] {
                curr[j] = curr[j] || prev[j-nums[i]]
            }
        }

        copy(prev, curr)
    }

    // fmt.Println(dp)
    return prev
}

func canPartition(nums []int) bool {
    totalSum, n := 0, len(nums)
    for i:=0; i<n; i++ {
        totalSum += nums[i]
    }

    if totalSum%2 != 0 {
        return false
    }

    res := tabulation1(nums, totalSum)
    return res[totalSum/2]
    
}