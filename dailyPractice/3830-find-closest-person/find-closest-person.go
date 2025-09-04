func findClosest(x int, y int, z int) int {
    abs := func(a, b int) int {
        if a>b {
            return a-b
        }
        return b-a
    }  

    xdist := abs(x, z)
    ydist := abs(y, z)

    if xdist < ydist {
        return 1
    } else if xdist > ydist {
        return 2
    }

    return 0
}