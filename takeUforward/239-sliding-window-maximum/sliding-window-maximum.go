func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    maxValIdxQueue := make([]int, 0, k)
    res := make([]int, n-k+1)

    for i:=0; i<k; i++ {
        for len(maxValIdxQueue)>0 && nums[maxValIdxQueue[len(maxValIdxQueue)-1]]<=nums[i] {
            maxValIdxQueue=maxValIdxQueue[:len(maxValIdxQueue)-1]
        }
        maxValIdxQueue = append(maxValIdxQueue, i)
    }

    res[0]=nums[maxValIdxQueue[0]]
    l, r := 0, k-1
    for r < n-1 {
        if maxValIdxQueue[0] == l {
            maxValIdxQueue=maxValIdxQueue[1:]
        }
        l++
        r++
        for len(maxValIdxQueue)>0 && nums[maxValIdxQueue[len(maxValIdxQueue)-1]]<=nums[r] {
            maxValIdxQueue=maxValIdxQueue[:len(maxValIdxQueue)-1]
        }
        maxValIdxQueue = append(maxValIdxQueue, r)
        res[l]=nums[maxValIdxQueue[0]]
    }

    return res
}