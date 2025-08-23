func subarraySum(nums []int, k int) int {
    m:=make(map[int]int)
    m[0]=1
    preSum, count := 0, 0
    for i:=0; i<len(nums); i++ {
        preSum+=nums[i]
        if val, exist := m[preSum-k]; exist {
            count+=val
        }
        m[preSum]++
    }

    return count
}