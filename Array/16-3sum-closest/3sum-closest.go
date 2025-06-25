func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n := len(nums)
    minDiff, res := 9999999, 0
    for i := 0; i < n; i++ {
        j, k := i+1, n-1
        for j < k {
            tempSum := nums[i]+nums[j]+nums[k]
            if tempSum == target {
                return target
            }
            if tempSum < target {
                if target - tempSum < minDiff {
                    minDiff = target - tempSum
                    res = tempSum
                }
                j++
            } else {
                if tempSum - target < minDiff {
                    minDiff = tempSum - target
                    res = tempSum
                }
                k--
            }
        }
    }

    return res
}