var res [][]int
func solve(nums, temp []int, vstd []bool) {
    if len(nums) == len(temp) {
        temp1 := make([]int, len(temp))
        copy(temp1, temp)
        res=append(res, temp1)
        return
    }

    for i:=0; i<len(vstd); i++ {
        if vstd[i] == false {
            // not take
            // solve(nums, temp, vstd)

            // take
            vstd[i]=true
            temp=append(temp, nums[i])
            solve(nums, temp, vstd)
            vstd[i]=false
            temp=temp[:len(temp)-1]
        }
    }
}

func permute(nums []int) [][]int {
    res=[][]int{}
    vstd:=make([]bool, len(nums))
    temp:=[]int{}

    solve(nums, temp, vstd)
    return res
}