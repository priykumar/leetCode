import "fmt"
func jump(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i:=0; i<n; i++ {
        dp[i]=99999
    }

    dp[0]=0
    for i:=0; i<n; i++ {
        for j:=i+1; j<=i+nums[i] && j<n; j++ {
            dp[j]=min(dp[j], dp[i]+1)
        }
        //fmt.Println(i, dp)
    }

    //fmt.Println(dp)
    return dp[n-1]
}