var res [][]int
func solve(nums []int, temp []int, currPos int) {
    if currPos == len(nums) {
        temp1 := make([]int, len(temp))
        copy(temp1, temp)
        res = append(res, temp1)
        return
    }

    // take
    temp=append(temp, nums[currPos])
    solve(nums, temp, currPos+1)
    temp=temp[:len(temp)-1]

    for currPos+1 < len(nums) && nums[currPos] == nums[currPos+1] {
        currPos++
    }
    // not take
    solve(nums, temp, currPos+1)
}

func subsetsWithDup(nums []int) [][]int {
    res = [][]int{}
    temp:=[]int{}
    sort.Ints(nums)
    solve(nums, temp, 0)

    return res
}