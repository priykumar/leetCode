func minimumArea(grid [][]int) int {
    rmin, rmax := -1, -1
    cmin, cmax := -1, -1
    r, c := len(grid), len(grid[0])

    // find min row
    found:=false
    for i:=0; i<r; i++ {
        for j:=0; j<c; j++ {
            if grid[i][j] == 1 {
                rmin=i
                found=true
                break
            }
        }
        if found {
            break
        }
    }

    // find max row
    found=false
    for i:=r-1; i>=0; i-- {
        for j:=0; j<c; j++ {
            if grid[i][j] == 1 {
                rmax=i
                found=true
                break
            }
        }
        if found {
            break
        }
    }

    // find min col
    found=false
    for i:=0; i<c; i++ {
        for j:=0; j<r; j++ {
            if grid[j][i] == 1 {
                cmin=i
                found=true
                break
            }
        }
        if found {
            break
        }
    }

    // find max row
    found=false
    for i:=c-1; i>=0; i-- {
        for j:=0; j<r; j++ {
            if grid[j][i] == 1 {
                cmax=i
                found=true
                break
            }
        }
        if found {
            break
        }
    }

    fmt.Println(rmin, rmax, cmin, cmax)
    return (rmax-rmin+1)*(cmax-cmin+1)
}