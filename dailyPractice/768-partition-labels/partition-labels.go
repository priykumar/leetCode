func check(set map[rune]int, totalCount [26]int) int {
    total := 0
    for k, v := range set {
        total+=v
        if totalCount[k-'a'] != v {
            return -1
        }
    }

    return total
}
func partitionLabels(str string) []int {
    set := make(map[rune]int)
    totalCount := [26]int{}
    res := []int{}

    for _, s := range str {
        totalCount[s-'a']++
    }

    for _, s := range str {
        set[s]++
        if val := check(set, totalCount); val != -1 {
            res=append(res, val)
            set=make(map[rune]int)
        }
    }

    return res
}