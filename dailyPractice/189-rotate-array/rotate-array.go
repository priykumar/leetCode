func rotate(nums []int, k int)  {
    n:=len(nums)
    k=k%n

    a := append(nums[n-k:], nums[:n-k]...)
    copy(nums, a)
    return 
}