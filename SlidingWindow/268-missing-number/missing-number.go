func missingNumber(nums []int) int {
    i := 0 
    for i<len(nums) {
        if i != nums[i] && nums[i] < len(nums){
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            i++
        }
    }
    
    for i:=0; i<len(nums); i++ {
        if i!=nums[i] {
            return i
        }
    }

    return len(nums)
}