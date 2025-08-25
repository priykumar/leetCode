func findDiagonalOrder(mat [][]int) []int {
    m, n := len(mat), len(mat[0])
    res := make([]int, 0, m*n)
    i, j := 0, 0
    upward := true

    for len(res) < m*n {
        res = append(res, mat[i][j])

        if upward {
            // moving up
            if j == n-1 {       // right border
                i++
                upward = false
            } else if i == 0 {  // top border
                j++
                upward = false
            } else {
                i--
                j++
            }
        } else {
            // moving down
            if i == m-1 {       // bottom border
                j++
                upward = true
            } else if j == 0 {  // left border
                i++
                upward = true
            } else {
                i++
                j--
            }
        }
    }
    return res
}
