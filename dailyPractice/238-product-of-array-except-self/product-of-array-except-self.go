func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    pre, suf := 1, 1
    for i:=0; i<n; i++ {
        res[i]=pre
        pre*=nums[i]
    }

    for i:=n-1; i>=0; i-- {
        res[i]*=suf
        suf*=nums[i]
    }

    return res
}