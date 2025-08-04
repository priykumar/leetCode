func isAnagram(s string, t string) bool {
    rune2Count := make(map[rune]int)
    for _, e := range s {
        rune2Count[e]++
    }

    for _, e := range t {
        rune2Count[e]--
        if rune2Count[e] == 0{
            delete(rune2Count, e)
        }
    }

    return len(rune2Count) == 0
}