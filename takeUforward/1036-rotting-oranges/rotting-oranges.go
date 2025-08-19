var dx []int = []int{1, -1, 0, 0}
var dy []int = []int{0, 0, 1, -1}
func orangesRotting(grid [][]int) int {
    type S struct {
        x int
        y int
    }

    queue := []S{}
    m, n := len(grid), len(grid[0])
    freshPresent := false
    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, S{i, j})
            } else if grid[i][j] ==1 {
                freshPresent = true
            }
        }
    }

    if !freshPresent {
        return 0
    }

    day := -1
    for len(queue) > 0 {
        fmt.Println(day, queue)
        day++
        sz := len(queue)
        for i:=0; i<sz; i++ {
            ux, uy := queue[i].x, queue[i].y
            for j:=0; j<4; j++ {
                vx, vy := ux+dx[j], uy+dy[j]
                if vx>=0 && vx<m && vy>=0 && vy<n && grid[vx][vy] == 1 {
                    grid[vx][vy] = 2
                    queue = append(queue, S{vx, vy})
                }
            }
        }
        queue=queue[sz:]
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }

    return day
}