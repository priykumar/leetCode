func solve(a, b string) string {
    m, n := len(a), len(b)
    i, j := 0, 0
    for i<m && j<n && a[i] == b[j] {
        i++
        j++
    }
    return a[:i]
}
func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }

    res:=solve(strs[0], strs[1])
    for i:=2; i<len(strs); i++ {
        res = solve(res, strs[i])
    }

    return res
}