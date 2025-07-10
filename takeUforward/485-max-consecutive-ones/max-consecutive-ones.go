func findMaxConsecutiveOnes(nums []int) int {
    i, j, n, maxlen := 0, 0, len(nums), 0

    for j<n {
        if nums[j] == 1 {
            if i == -1 {
                i=j
            }
        } else {
            if i != -1 {
                maxlen = max(maxlen, j-i)
            }
            i = -1
        }
        j++
    }

    if i!=-1 {
        maxlen = max(maxlen, j-i)
    }

    return maxlen
}