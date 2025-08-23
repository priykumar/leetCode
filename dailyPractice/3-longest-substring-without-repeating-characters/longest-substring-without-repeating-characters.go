func lengthOfLongestSubstring(s string) int {
    char := make(map[byte]int)
    uniqueCharCount := 0
    start := 0
    res:=0
    for i:=0; i<len(s); i++ {
        if val, exist := char[s[i]]; exist {
            currStart := val+1
            start=max(currStart, start)
            uniqueCharCount=i-start+1
        } else {
            uniqueCharCount++
        }
        res = max(res, uniqueCharCount)
        char[s[i]]=i
    }
    return res
}