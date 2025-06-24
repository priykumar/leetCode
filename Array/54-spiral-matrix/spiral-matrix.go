func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
    top, btm, left, right := 0, m-1, 0, n-1
    res := make([]int, m*n)

    run, idx := 0, -1
    for top <= btm && left <= right {
        if run%4 == 0 {
            for i:=left; i<=right; i++ {
                idx++
                res[idx] = matrix[top][i]
            }
            top++
        } else if run%4==1 {
            for i:=top; i<=btm; i++ {
                idx++
                res[idx] = matrix[i][right]
            }
            right--
        } else if run%4==2 {
            for i:=right; i>=left; i-- {
                idx++
                res[idx] = matrix[btm][i]
            }
            btm--
        } else {
            for i:=btm; i>=top; i-- {
                idx++
                res[idx] = matrix[i][left]
            }
            left++
        }

        run++
    }

    return res
}