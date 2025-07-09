func majorityElement(nums []int) []int {
    n := len(nums)
    if n==1 {
        return nums
    } else if n==2 {
        if nums[0] == nums[1] {
            return []int{nums[0]}
        } else {
            return nums
        }
    }

    sort.Ints(nums)
    limit, i := n/3, 0
    currIdx := 0
    res:=[]int{}
    for currIdx + limit < n {
        if nums[currIdx] == nums[currIdx+limit] {
            res=append(res, nums[currIdx])
            i=currIdx+limit
            for i<n-1 && nums[i+1] == nums[i] {
                i++
            }
            currIdx = i+1
        } else {
            currIdx++
        }
    }

    return res
}