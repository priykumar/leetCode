func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

    result:=[][]int{}
    l, r := intervals[0][0], intervals[0][1]
    for i:=1; i<len(intervals); i++ {
        if intervals[i][0] <= r {
            r = max(r, intervals[i][1])
        } else {
            result = append(result, []int{l, r})
            l, r = intervals[i][0], intervals[i][1]
        }
    }

    result = append(result, []int{l, r})

    return result
}