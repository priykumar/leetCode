func countDays(days int, meetings [][]int) int {
    sort.SliceStable(meetings, func(i, j int)bool {
        return meetings[i][0] < meetings[j][0]
    })

    count:=0
    n := len(meetings)
    l, r := meetings[0][0], meetings[0][1]
    count+=(l-1)
    i := 1
    for i<n {
        if meetings[i][0] <= r {
            r=max(r, meetings[i][1])
        } else {
            count+=(meetings[i][0]-(r+1))
            l, r = meetings[i][0], meetings[i][1]
        }
        i++
    }
    count+=(days-r)
    return count
}