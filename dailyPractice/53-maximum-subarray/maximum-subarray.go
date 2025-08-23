func maxSubArray(nums []int) int {
    // Kadane's
    sum:=0
    res := nums[0]
    for i:=0; i<len(nums); i++ {
        sum+=nums[i]
        res = max(res, sum)
        if sum < 0 {
            sum=0
        }
    }

    return res
}