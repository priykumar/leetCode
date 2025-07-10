func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    i, n := 0, len(nums)
    result:=[][]int{}
    for i < n-2 {
        j, k := i+1, n-1
        for j < k {
            if nums[i] + nums[j] + nums[k] == 0 {
                result=append(result, []int{nums[i], nums[j], nums[k]})
                curr_j, curr_k := j, k
                for j < k && nums[curr_j] == nums[j] {
                    j++
                }
                for j < k && nums[curr_k] == nums[k] {
                    k--
                }
            } else if nums[i] + nums[j] + nums[k] < 0 {
                j++
            } else {
                k--
            }
        }

        curr_i := i
        for i < n-2 && nums[curr_i] == nums[i] {
            i++
        }
    }

    return result
}