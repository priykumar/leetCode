func longestSubarray(nums []int) int {
    pp0, p0 := -1, -1
    n:=len(nums)
    length := 0
    zeroCount := 0
    for i:=0; i<n; i++ {
        if nums[i] == 0 {
            length = max(length, i-1-pp0)
            pp0, p0 = p0, i
            if nums[i] == 0 {
                zeroCount++
            }
            fmt.Println(i, pp0, p0, length)
        } else if i == n-1 {
            length = max(length, i-pp0)
            fmt.Println(i, pp0, p0, length)
        }
    }

    if zeroCount  < 2{
        return n-1
    }

    return length-1
}