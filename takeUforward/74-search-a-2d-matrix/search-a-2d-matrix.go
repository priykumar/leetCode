func findRow(matrix [][]int, target int) int {
    l, r := 0, len(matrix)-1
    if target >= matrix[r][0] {
        return r
    }

    fmt.Println(l, r)
    res:=-1
    for l<=r {
        m:=l+(r-l)/2
        if matrix[m][0]<=target {
            res = m
        }
        if target < matrix[m][0] {
            r=m-1
        } else {
            l=m+1
        }
    }

    return res
}


func searchMatrix(matrix [][]int, target int) bool {
    row := findRow(matrix, target)
    if row == -1 {
        return false
    }
    
    l, r := 0, len(matrix[0])-1
    fmt.Println(row)
    for l<=r {
        m := l+(r-l)/2
        if matrix[row][m] == target {
            return true
        }
        if target < matrix[row][m] {
            r=m-1
        } else {
            l=m+1
        }
    }
    return false
}