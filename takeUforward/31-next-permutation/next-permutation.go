func nextPermutation(nums []int)  {
    n := len(nums)
    if n == 1 {
        return
    }
    
    i := n-2
    for i >= 0 && nums[i] >= nums[i+1]{
        i--
    }

    if i < 0 {
        sort.Ints(nums)
        return 
    }

    // find next largest element to nums[i]
    minDiff:=1000
    eleIdx:=-1
    for j:=i+1; j<n && nums[j]>nums[i]; j++ {
        if nums[j]-nums[i] < minDiff {
            minDiff = nums[j]-nums[i]
            eleIdx=j
        }
    }

    nums[i], nums[eleIdx] = nums[eleIdx], nums[i]
    sort.Ints(nums[i+1:])
}