func canPartition(nums []int) bool {
    totalSum := 0
    for i := range nums {
        totalSum+=nums[i]
    }
    if totalSum%2!=0 {
        return false
    }

    prev, curr := make([]bool, (totalSum/2)+1), make([]bool, (totalSum/2)+1)
    prev[0]=true
    if nums[0] < len(prev) {
        prev[nums[0]] = true
    }

    for i:=1; i<len(nums); i++ {
        // fmt.Println(prev)
        curr[0]=true
        for j:=0; j<(totalSum/2)+1; j++ {
            // not take
            curr[j]=prev[j]
            // take
            if j>=nums[i] {
                curr[j] = curr[j] || prev[j-nums[i]]
            }
        }
        copy(prev, curr)
    }
    // fmt.Println()
    // fmt.Println(prev)
    return prev[len(prev)-1]
}