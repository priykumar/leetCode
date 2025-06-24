func subarraySum(nums []int, k int) int {
    pSum, count := 0, 0
    pSumFreq:=map[int]int{}
    pSumFreq[0]=1
    for i:=0; i<len(nums); i++ {
        pSum+=nums[i]
        count+=pSumFreq[pSum-k]
        pSumFreq[pSum]++
    }
    return count
}