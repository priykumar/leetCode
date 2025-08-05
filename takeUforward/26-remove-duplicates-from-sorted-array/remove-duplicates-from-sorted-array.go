func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 1
    }

    l, r := 0, 1
    for r < n {
        if nums[r] == nums[l] {
            r++
        } else {
            l++
            nums[l]=nums[r]
            r++
        }
    }

    return l+1
}