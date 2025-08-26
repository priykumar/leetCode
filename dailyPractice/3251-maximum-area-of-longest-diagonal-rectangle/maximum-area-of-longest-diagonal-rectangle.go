func areaOfMaxDiagonal(dimensions [][]int) int {
    maxArea, maxDiag := 0, 0
    for _, d := range dimensions {
        tempDiag := d[0]*d[0] + d[1]*d[1]
        if tempDiag > maxDiag {
            maxArea = d[0]*d[1]
            maxDiag = tempDiag
        } else if tempDiag == maxDiag {
            maxArea = max(maxArea, d[0]*d[1])
        }
    }

    return maxArea
}