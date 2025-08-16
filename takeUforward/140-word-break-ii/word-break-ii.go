var dict map[string]bool
var res []string

func solve(start int, s string, tempRes string) {
    if start >= len(s) {
        res = append(res, tempRes[1:]) // remove 1 leading space
        return 
    }

    for j:=start; j<len(s); j++ {
        if dict[s[start:j+1]] {
            solve(j+1, s, tempRes+ " " + s[start:j+1])
        }
    }
}

func wordBreak(s string, wordDict []string) []string {
    dict = make(map[string]bool)
    res = make([]string, 0)
    for _, w := range wordDict {
        dict[w] = true
    }

    tempRes := ""
    solve(0, s, tempRes)
    return res
}