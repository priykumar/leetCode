func separateDigits(nums []int) []int {
    n:=len(nums)
    res := []int{}
    for i:=n-1; i>=0; i-- {
        curr := nums[i]
        for curr > 0{
            res=append(res, curr%10)
            curr=curr/10
        }
    }

    n=len(res)
    for i, j :=0, n-1; i<j; i, j=i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

    return res
}