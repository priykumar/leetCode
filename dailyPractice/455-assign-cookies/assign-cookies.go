func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    m, n := len(g), len(s)
    i, j := 0, 0
    for i<m && j<n {
        if s[j] >= g[i] {
            i++
            j++
        } else {
            j++
        }
    }

    return i
}