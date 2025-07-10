func removeDuplicates(nums []int) int {
    currIdx, fwdIdx, n := 0, 1, len(nums)
    for fwdIdx<n {
        for fwdIdx<n && nums[fwdIdx]==nums[fwdIdx-1] {
            fwdIdx++
        }

        if fwdIdx<n {
            currIdx++
            nums[currIdx]=nums[fwdIdx]   
            fwdIdx++
        }
    }

    return currIdx+1
}