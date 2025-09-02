var xmap, ymap map[int][][]int
func pointslieinbetween(xmax, xmin, ymax, ymin int) bool{
    // fmt.Println("------------here")
    // fmt.Println(xmax, xmin, ymax, ymin)
    for i:=xmin; i<=xmax; i++ {
        if val, exist := xmap[i]; exist {
            for _, v := range val {
                if (v[0]==xmin && v[1]==ymax) || (v[0]==xmax && v[1]==ymin) {
                    continue
                }
                // fmt.Println(v)
                if ymin<=v[1] && v[1]<=ymax{
                    // fmt.Println("true")
                    return true
                }
            }
        }
    }

    // fmt.Println("false")
    return false
}

func numberOfPairs(points [][]int) int {
    xmap, ymap = make(map[int][][]int), make(map[int][][]int)
    for _, p := range points {
        if _, exist := xmap[p[0]]; !exist {
            xmap[p[0]] = make([][]int, 0)
        }
        xmap[p[0]] = append(xmap[p[0]], p)
    }

    for _, p := range points {
        if _, exist := ymap[p[1]]; !exist {
            ymap[p[1]] = make([][]int, 0)
        }
        ymap[p[1]] = append(ymap[p[1]], p)
    }

    fmt.Println(xmap, ymap)
    count := 0
    n := len(points)
    for i:=0; i<n-1; i++ {
        for j:=i+1; j<n; j++ {
            xi, yi, xj, yj := points[i][0], points[i][1], points[j][0], points[j][1]
            if (xi<=xj && yi>=yj) {
                // fmt.Println("h1", xi,yi,xj,yj)
                if !pointslieinbetween(xj, xi, yi, yj){
                    count++
                    fmt.Println(count)
                }
            } else if (xj<=xi && yj>=yi) {
                // fmt.Println("h2", xi,yi,xj,yj)
                if !pointslieinbetween(xi, xj, yj, yi) {
                    count++
                    fmt.Println(count)
                }
            }
        }
    }

    return count
}