var dx []int = []int{1, 0, -1, 0}
var dy []int = []int{0, 1, 0, -1}
func bfs(grid [][]byte, x, y int) {
    type S struct {
        x int
        y int
    }

    grid[x][y] = "2"[0]
    m, n := len(grid), len(grid[0])
    q := []S{{x,y}}
    for len(q) > 0{
        ux, uy := q[0].x, q[0].y
        q=q[1:]

        for i:=0; i<4; i++ {
            vx, vy := ux+dx[i], uy+dy[i]
            if vx>=0 && vy>=0 && vx<m && vy<n && grid[vx][vy] == "1"[0] {
                grid[vx][vy] = "2"[0]
                // fmt.Println(x,y,vx,vy)
                q=append(q, S{vx, vy})
            }
        }
    }
}

func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    count:=0
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == "1"[0] {
                bfs(grid, i, j)
                // fmt.Println(grid)
                count++
            }
        }
    }

    return count
}