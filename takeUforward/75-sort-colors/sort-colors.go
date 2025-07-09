func sortColors(nums []int)  {
    l, m, h := 0, 0, len(nums)-1
    
    for m<=h{
        if nums[m] == 0 {
            nums[l], nums[m] = nums[m], nums[l]
            l++
            m++
        } else if nums[m] == 1 {
            m++
        } else {
            nums[h], nums[m] = nums[m], nums[h]
            h--
        }
    }

}