func findDuplicate(nums []int) int {
    i:=0
    for i<len(nums) {
        if i+1 != nums[i] && nums[nums[i]-1] >= 0 && nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        } else {
            i++
        }
    }

    for i:=1; i<=len(nums); i++ {
        if i!=nums[i-1] {
            return nums[i-1]
        }
    }

    return -1
}