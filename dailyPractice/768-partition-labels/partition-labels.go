func check(currCount, totalCount [26]int) int {
    total := 0
    for i:=0; i<26; i++ {
        total+=currCount[i]
        if currCount[i] != 0 && currCount[i] != totalCount[i] {
            return -1
        }
    }

    return total
}
func partitionLabels(str string) []int {
    totalCount := [26]int{}
    currCount := [26]int{}
    res := []int{}

    for _, s := range str {
        totalCount[s-'a']++
    }

    for _, s := range str {
        currCount[s-'a']++
        if currCount[s-'a'] == totalCount[s-'a'] {
            if val := check(currCount, totalCount); val != -1 {
                res=append(res, val)
                currCount=[26]int{}
            }
        }
    }

    return res
}