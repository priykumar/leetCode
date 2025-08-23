func merge(intervals [][]int) [][]int {
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    fmt.Println(intervals)
    l, r := intervals[0][0], intervals[0][1]
    i, n := 1, len(intervals)
    res := make([][]int, 0)
    for i < n {
        if intervals[i][0] <= r {
            r = max(r, intervals[i][1])
        } else {
            res = append(res, []int{l, r})
            l, r = intervals[i][0], intervals[i][1]
        }
        i++
    }
    res = append(res, []int{l, r})

    return res
}