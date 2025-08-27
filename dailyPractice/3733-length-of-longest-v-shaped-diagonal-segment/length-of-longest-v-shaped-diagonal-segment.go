func lenOfVDiagonal(grid [][]int) int {
    n, m := len(grid), len(grid[0])

    downRight := make([][]int, n)
    for i := range n {
        downRight[i] = make([]int, m)
    }
    for i := range n {
        if grid[i][0] != 1 {
            downRight[i][0] = 1
        }
    }
    for i := range m {
        if grid[0][i] != 1 {
            downRight[0][i] = 1
        }
    }
    for r := 1; r < n; r++ {
        for c := 1; c < m; c++ {
            if grid[r][c] == 0 || grid[r][c] == 2 {
                if grid[r][c] == -grid[r-1][c-1] + 2 {
                    downRight[r][c] = 1 + downRight[r-1][c-1]
                } else {
                    downRight[r][c] = 1
                }
            }
        }
    }

    downLeft := make([][]int, n)
    for i := range n {
        downLeft[i] = make([]int, m)
    }
    for i := range n {
        if grid[i][m-1] != 1 {
            downLeft[i][m-1] = 1
        }
    }
    for i := range m {
        if grid[0][i] != 1 {
            downLeft[0][i] = 1
        }
    }
    for r := 1; r < n; r++ {
        for c := m-2; c >= 0; c-- {
            if grid[r][c] == 0 || grid[r][c] == 2 {
                if grid[r][c] == -grid[r-1][c+1] + 2 {
                    downLeft[r][c] = 1 + downLeft[r-1][c+1]
                } else {
                    downLeft[r][c] = 1
                }
            }
        }
    }

    upLeft := make([][]int, n)
    for i := range n {
        upLeft[i] = make([]int, m)
    }
    for i := range n {
        if grid[i][m-1] != 1 {
            upLeft[i][m-1] = 1
        }
    }
    for i := range m {
        if grid[n-1][i] != 1 {
            upLeft[n-1][i] = 1
        }
    }
    for r := n-2; r >= 0; r-- {
        for c := m-2; c >= 0; c-- {
            if grid[r][c] == 0 || grid[r][c] == 2 {
                if grid[r][c] == -grid[r+1][c+1] + 2 {
                    upLeft[r][c] = 1 + upLeft[r+1][c+1]
                } else {
                    upLeft[r][c] = 1
                }
            }
        }
    }

    upRight := make([][]int, n)
    for i := range n {
        upRight[i] = make([]int, m)
    }
    for i := range n {
        if grid[i][0] != 1 {
            upRight[i][0] = 1
        }
    }
    for i := range m {
        if grid[n-1][i] != 1 {
            upRight[n-1][i] = 1
        }
    }
    for r := n-2; r >= 0; r-- {
        for c := 1; c < m; c++ {
            if grid[r][c] == 0 || grid[r][c] == 2 {
                if grid[r][c] == -grid[r+1][c-1] + 2 {
                    upRight[r][c] = 1 + upRight[r+1][c-1]
                } else {
                    upRight[r][c] = 1
                }
            }
        }
    }

    out := 0
    for r := range n {
        for c := range m {
            if grid[r][c] == 1 {
                out = max(out, 1)

                // UP LEFT
                dir := []int{-1, -1}
                expected := 2
                length := 0
                row, col := r+dir[0], c+dir[1]
                for row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == expected {
                    length++
                    expected = -expected + 2
                    bendr, bendc := row-1, col+1
                    tail := 0
                    if bendr >= 0 && bendr < n && bendc >= 0 && bendc < m && grid[bendr][bendc] == expected {
                        tail = downLeft[bendr][bendc]
                    }
                    out = max(out, 1 + length + tail)
                    row += dir[0]
                    col += dir[1]
                }

                // UP RIGHT
                dir = []int{-1, 1}
                expected = 2
                length = 0
                row, col = r+dir[0], c+dir[1]
                for row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == expected {
                    length++
                    expected = -expected + 2
                    bendr, bendc := row+1, col+1
                    tail := 0
                    if bendr >= 0 && bendr < n && bendc >= 0 && bendc < m && grid[bendr][bendc] == expected {
                        tail = upLeft[bendr][bendc]
                    }
                    out = max(out, 1 + length + tail)
                    row += dir[0]
                    col += dir[1]
                }

                // DOWN RIGHT
                dir = []int{1, 1}
                expected = 2
                length = 0
                row, col = r+dir[0], c+dir[1]
                for row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == expected {
                    length++
                    expected = -expected + 2
                    bendr, bendc := row+1, col-1
                    tail := 0
                    if bendr >= 0 && bendr < n && bendc >= 0 && bendc < m && grid[bendr][bendc] == expected {
                        tail = upRight[bendr][bendc]
                    }
                    out = max(out, 1 + length + tail)
                    row += dir[0]
                    col += dir[1]
                }

                // DOWN LEFT
                dir = []int{1, -1}
                expected = 2
                length = 0
                row, col = r+dir[0], c+dir[1]
                for row >= 0 && row < n && col >= 0 && col < m && grid[row][col] == expected {
                    length++
                    expected = -expected + 2
                    bendr, bendc := row-1, col-1
                    tail := 0
                    if bendr >= 0 && bendr < n && bendc >= 0 && bendc < m && grid[bendr][bendc] == expected {
                        tail = downRight[bendr][bendc]
                    }
                    out = max(out, 1 + length + tail)
                    row += dir[0]
                    col += dir[1]
                }
            }
        }
    }
    
    return out
}