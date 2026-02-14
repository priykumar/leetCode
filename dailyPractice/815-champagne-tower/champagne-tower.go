func getVal(a float64) float64 {
    return max((a-1)/2, 0.0)
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
    prev, curr := make([]float64, 100), make([]float64, 100)
    prev[0] = float64(poured)

    for i:=1; i<=query_row; i++ {
        curr[0] = getVal(prev[0])
        curr[i] = getVal(prev[i-1])
        for j:=1; j<i; j++ {
            curr[j]=getVal(prev[j-1])+getVal(prev[j])
        }

        copy(prev, curr)
    }

    return min(prev[query_glass], 1.0)
}