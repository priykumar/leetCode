func generate(numRows int) [][]int {
    res:=[][]int{{1}}
    if numRows == 1 {
        return res
    }
    res=append(res, []int{1,1})
    
    for i:=2; i<numRows; i++ {
        temp:=make([]int, i+1)
        temp[0], temp[i] = 1, 1
        for j:=1; j<i; j++ {
            temp[j]=res[i-1][j-1] + res[i-1][j]
        }
        res=append(res, temp)
    }

    return res
}