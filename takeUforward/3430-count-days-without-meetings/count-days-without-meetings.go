func countDays(days int, meetings [][]int) int {
    sort.SliceStable(meetings, func(i, j int)bool {
        return meetings[i][0] < meetings[j][0]
    })

    count:=0
    fmt.Println(count)

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

    // maxEnd := meetings[0][1]
    // for i:=1; i<len(meetings); i++ {
    //     maxEnd = max(maxEnd, meetings[i][1])
    //     if meetings[i-1][1]+1<meetings[i][0] {
    //         count+=(meetings[i][0]-(meetings[i-1][1]+1))
    //     }
    // }
    fmt.Println(count)
    count+=(days-r)
    fmt.Println(count)
    fmt.Println(meetings)
    return count
}

// [1,3] [2,4]