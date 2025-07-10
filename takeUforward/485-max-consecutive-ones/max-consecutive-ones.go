func findMaxConsecutiveOnes(nums []int) int {
    start, curr, n, maxlen := 0, 0, len(nums), 0

    for curr<n {
        if nums[curr] == 1 {
            if start == -1 {
                start=curr
            }
        } else {
            if start != -1 {
                maxlen = max(maxlen, curr-start)
            }
            start = -1
        }
        curr++
    }

    if start!=-1 {
        maxlen = max(maxlen, curr-start)
    }

    return maxlen
}