func check(board [][]byte, r, c int) bool {
    for i:=0; i<9; i++ {
        if i == r {
            continue
        }
        if board[i][c] == board[r][c] {
            return false
        }
    }

    for i:=0; i<9; i++ {
        if i == c {
            continue
        }
        if board[r][i] == board[r][c] {
            return false
        }
    }

    offset_r := r/3
    offset_c := c/3
    for i:=0; i<3; i++ {
        for j:=0; j<3; j++ {
            x, y := i+offset_r*3, j+offset_c*3
            if x == r && y == c {
                continue
            }
            if board[r][c] == board[x][y] {
                return false
            }
        }
    }

    return true
}

func solve(board *[][]byte, r, c int) bool {
    if c == 9{
        r+=1
        c=0
    }
    if r >= 9 {
        return true
    }

    if (*board)[r][c] == '.' {
        for num:=1; num<=9; num++ {
            (*board)[r][c] = byte(num+'0')
            // fmt.Println("here1111")
            if check(*board, r, c) {
                // fmt.Println("here2222")
                if solve(board, r, c+1) {
                    // fmt.Println("here")
                    return true
                }
            }
            (*board)[r][c] = '.'
        }
    } else {
        return solve(board, r, c+1)
    }

    return false
}

func solveSudoku(board [][]byte)  {
    solve(&board, 0, 0)
}