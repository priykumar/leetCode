import "fmt"

var res [][]int
func backtrack1(candidates, currArr []int, target, idx, sum int) {
    if sum == target {
        a:=make([]int, len(currArr))
        copy(a, currArr)
        res=append(res, a)
        return
    }

    if idx >= len(candidates) || sum > target{
        return
    }

    // for i:=idx; i<len(candidates); i++ {
        // not take
        backtrack1(candidates, currArr, target, idx+1, sum)

        // take
        currArr=append(currArr, candidates[idx])
        backtrack1(candidates, currArr, target, idx, sum+candidates[idx])
        currArr=currArr[:len(currArr)-1]
    // }
}

func combinationSum(candidates []int, target int) [][]int {
    res=[][]int{}
    currArr:=[]int{}
    backtrack1(candidates, currArr, target, 0, 0)
    return res
}