var res [][]int
func solve(candidates, temp []int, target, currPos, currSum int) {
    if currSum == target {
        temp1:=make([]int, len(temp))
        copy(temp1, temp)
        res=append(res, temp1)
        return
    }

    if currPos >= len(candidates) || currSum > target {
        return
    }

    // take
    temp=append(temp, candidates[currPos])
    solve(candidates, temp, target, currPos+1, currSum+candidates[currPos])
    temp=temp[:len(temp)-1]

    for currPos+1 < len(candidates) && candidates[currPos] == candidates[currPos+1] {
        currPos+=1
    }
    // not take
    solve(candidates, temp, target, currPos+1, currSum)
}

func combinationSum2(candidates []int, target int) [][]int {
    res = [][]int{}
    temp := []int{}
    sort.Ints(candidates)
    solve(candidates, temp, target, 0, 0)

    return res
}