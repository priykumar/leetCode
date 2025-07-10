func search(nums []int, target int) int {
    first, last := nums[0], nums[len(nums)-1]
    l, r := 0, len(nums)-1

    for l<=r {
        m:=l+(r-l)/2
        if nums[m]==target {
            return m
        }

        if nums[m]<first {
            if target<nums[m] || target>last {
                r=m-1
            } else {
                l=m+1
            }
        } else {
            if target>nums[m] || target<first {
                l=m+1
            } else {
                r=m-1
            }
        }
    }

    return -1
}