func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    res[0] = nums[0]
    for i:=1; i<n; i++ {
        res[i]=nums[i]*res[i-1]
    }

    postProduct := nums[n-1]
    res[n-1] = res[n-2]
    for i:=n-2; i>0; i-- {
        res[i] = res[i-1]*postProduct
        postProduct=postProduct*nums[i]
    }
    res[0]=postProduct

    return res
}