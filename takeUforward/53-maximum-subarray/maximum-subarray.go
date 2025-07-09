func maxSubArray(nums []int) int {
    maxSum:=int(-1e9)
    sum:=0
    for i:=0; i<len(nums); i++ {
        sum+=nums[i]
        maxSum=max(maxSum, sum)

        if sum<0 {
            sum=0
        }
    }

    return maxSum
}