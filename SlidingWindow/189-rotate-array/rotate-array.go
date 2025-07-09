func reverse(nums []int) {
    i, j := 0, len(nums)-1
    for i<j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
func rotate(nums []int, k int)  {
    k=k%len(nums)
    if k==0 {
        return 
    }
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}