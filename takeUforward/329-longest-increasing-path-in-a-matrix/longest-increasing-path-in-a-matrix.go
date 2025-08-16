var dx []int = []int{-1, 0, 1, 0}
var dy []int = []int{0, -1, 0, 1}

var res int
var dp [][]int
func dfs(matrix [][]int, x, y, m, n int, currLen int) {
    if dp[x][y] != 0 {
        res = max(res, currLen + dp[x][y] - 1)
        return
    }

    res = max(res, currLen)

    origVal := matrix[x][y]
    matrix[x][y] = -1 // visited

    best := 1
    for i:=0; i<4; i++ {
        vx, vy := x+dx[i], y+dy[i]
        if vx>=0 && vy>=0 && vx<m && vy<n && matrix[vx][vy] > origVal && matrix[vx][vy] != -1 { 
            dfs(matrix, vx, vy, m, n, currLen+1)
            best = max(best, 1+dp[vx][vy])
        }
    }
    matrix[x][y] = origVal

    dp[x][y] = best
    res = max(res, currLen)
}

func longestIncreasingPath(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    res = 0

    dp = make([][]int, m)
    for i:=0; i<m; i++ {
        dp[i] = make([]int, n)
    }
    
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dfs(matrix, i, j, m, n, 1)
        }
    }

    return res
}