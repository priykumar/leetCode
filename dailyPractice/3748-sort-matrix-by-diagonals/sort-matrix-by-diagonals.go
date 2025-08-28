func sortMatrix(grid [][]int) [][]int {
    tempArr := make([]int, len(grid))

    // Sort bottom left in DSC order
    i := 0
    n:=len(grid)
    for i<n {
        k:=0
        for ii, jj:=i, 0; ii<n && jj<n-i; ii, jj = ii+1, jj+1 {
            tempArr[k]=grid[ii][jj]
            k++
        }

        sort.Ints(tempArr[:k])
        k=k-1
        for ii, jj:=i, 0; ii<n && jj<n-i; ii, jj = ii+1, jj+1 {
            grid[ii][jj] = tempArr[k]
            k--
        }
        i++
    }

    // fmt.Println("---")
    tempArr = make([]int, len(grid))
    // Sort top right in ASC order
    j := 1
    for j<n {
        k:=0
        for ii, jj:=0, j; ii<n-j && jj<n; ii, jj = ii+1, jj+1 {
            tempArr[k]=grid[ii][jj]
            k++
        }

        sort.Ints(tempArr[:k])
        k=0
        for ii, jj:=0, j; ii<n-j && jj<n; ii, jj = ii+1, jj+1 {
            grid[ii][jj] = tempArr[k]
            k++
        }
        j++
    }

    return grid
}