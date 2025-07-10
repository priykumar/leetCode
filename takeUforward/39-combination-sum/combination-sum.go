var res [][]int
func solve(candidates, temp []int, target, currPos, currSum int) {
    if currSum == target {
        temp1 := make([]int, len(temp))
        copy(temp1, temp)
        res = append(res, temp1)
        fmt.Println(res)
        return
    }

    if currPos >= len(candidates) || currSum > target {
        return 
    }

    // not take
    solve(candidates, temp, target, currPos+1, currSum)

    // take
    temp=append(temp, candidates[currPos])
    solve(candidates, temp, target, currPos, currSum+candidates[currPos])
    temp=temp[:len(temp)-1]
}

func combinationSum(candidates []int, target int) [][]int {
    res = [][]int{}
    temp := []int{}
    solve(candidates, temp, target, 0, 0)
    // fmt.Println(res)
    return res
}