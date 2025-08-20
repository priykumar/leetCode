func countSquares(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    count := 0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if i==0 || j==0 {
                if matrix[i][j] == 1 {
                    count++
                }
                continue
            }
            if matrix[i][j] == 1 {
                matrix[i][j] = 1+min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1]))
                count+=matrix[i][j]
            }
        }
    }
    return count
}