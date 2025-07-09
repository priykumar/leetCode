func longestOnes(nums []int, k int) int {
    zeroCount, n := 0, len(nums)

    i, j := 0, 0
    maxlen := 0
    for i<n && j<n {
        if nums[j] == 0 {
            zeroCount++
            if zeroCount > k {
                for i<j && nums[i]!= 0 {
                    i++
                }

                if nums[i] == 0 {
                    i++
                    zeroCount--
                }
            }
        }

        if zeroCount <= k {
            maxlen = max(maxlen, j-i+1)
        }

        j++
    }

    return maxlen
}