func findKDistantIndices(nums []int, key int, k int) []int {
    n:=len(nums)
    res:=[]int{}
    for i:=0; i<n; i++ {
        if nums[i]==key {
            start:=max(0, i-k)
            end:=min(n-1, i+k)

            if len(res) != 0 {
                start=max(start, res[len(res)-1]+1)
            }
            for j:=start; j<=end; j++ {
                res=append(res, j)
            }
        }
    }
    return res
}