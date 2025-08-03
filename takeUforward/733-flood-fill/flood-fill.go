

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    dx:=[]int{1, 0, -1, 0}
    dy:=[]int{0, 1, 0, -1}

    q := [][]int{}
    q = append(q, []int{sr, sc})
    
    m, n := len(image), len(image[0])
    origColor := image[sr][sc]
    if origColor == color {
        return image
    }
    image[sr][sc] = color

    for len(q) > 0{
        fmt.Println("here1")
        ux, uy := q[0][0], q[0][1]
        q=q[1:]

        for i:=0; i<4; i++ {
            fmt.Println("here2")
            vx, vy := ux+dx[i], uy+dy[i]
            if 0 <= vx && vx < m && 0 <= vy && vy < n && image[vx][vy] == origColor {
                fmt.Println("here3")
                image[vx][vy] = color
                q = append(q, []int{vx, vy})
            }
        }
    }

    return image
}