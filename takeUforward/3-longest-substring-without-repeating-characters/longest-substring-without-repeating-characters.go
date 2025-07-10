func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n <= 1 {
        return n
    }

    isCharPresent := map[byte]int{}
    maxlen:=0
    i, j := 0, 0
    for j < n {
        if isCharPresent[s[j]] == 0 {
            maxlen = max(maxlen, j-i+1)
            isCharPresent[s[j]]++
            j++
        } else {
            isCharPresent[s[i]]--
            i++
        }

    }

    return maxlen
}