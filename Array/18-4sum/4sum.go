func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n:=len(nums)
    if n < 4 {
        return [][]int{}
    }

    res:=[][]int{}
    a:=0
    for a < n-3 {
        b:=a+1
        for b < n-2 {
            c, d := b+1, n-1
            for c < d {
                if nums[a]+nums[b]+nums[c]+nums[d] == target {
                    res=append(res, []int{nums[a], nums[b], nums[c], nums[d]})
                    curr_c, curr_d := c, d
                    for c < d && nums[curr_c] == nums[c] {
                        c++
                    }
                    for c < d && nums[curr_d] == nums[d] {
                        d--
                    }
                } else if nums[a]+nums[b]+nums[c]+nums[d] < target {
                    c++
                } else {
                    d--
                }
            }

            curr_b := b
            for b < n-2 && nums[curr_b] == nums[b] {
                b++
            }
        }

        curr_a := a
        for a < n-3 && nums[curr_a] == nums[a] {
            a++
        }
    }

    return res
}