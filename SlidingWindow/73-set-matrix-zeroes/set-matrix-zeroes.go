func setZeroes(matrix [][]int)  {
    rows, cols := map[int]bool{}, map[int]bool{}
    m, n := len(matrix), len(matrix[0])
    
    // make list of rows and cols with zeros
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if matrix[i][j] == 0 {
                rows[i]=true
                cols[j]=true
            }
        }
    }

    // populate matrix
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if rows[i] || cols[j] {
                matrix[i][j]=0
            }
        }
    }

    //return matrix
}