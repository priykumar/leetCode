func isValid(col []int, x, y int) bool {
    for i:=0; i<x; i++ {
        r, c := i, col[i]
        if x == r || y == c || x+y == r+c || x-y == r-c {
            return false
        }
    }

    return true
}

func solve(n, currRow int, col []int, res *[][]string) {
    if currRow == n {
        board := []string{}
		for i := 0; i < n; i++ {
			rowStr := ""
			for j := 0; j < n; j++ {
				if col[i] == j {
					rowStr += "Q"
				} else {
					rowStr += "."
				}
			}
			board = append(board, rowStr)
		}
		*res = append(*res, board)
		return
    }

    for c:=0; c<n; c++ {
        if isValid(col, currRow, c) {
            col[currRow] = c
            solve(n, currRow+1, col, res) 
            col[currRow] = -1
        }
    }

    return
}

func solveNQueens(n int) [][]string {
    if n==1 {
        return [][]string{{"Q"}}
    }

    res := [][]string{}
	col := make([]int, n)
	for i := range col {
		col[i] = -1
	}
	solve(n, 0, col, &res)
	return res
}