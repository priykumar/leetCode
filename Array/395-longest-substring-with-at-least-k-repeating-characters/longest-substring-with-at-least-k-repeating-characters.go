func longestSubstring(s string, k int) int {
    if len(s) < k {
		return 0
	}

    if k==1 {
        return len(s)
    }

    // Freq of each chars in the string
    m := map[rune]int{}
    for _, e := range s {
        m[e]++
    }

    // Check if any char with less than 'k' count
    invalidFound:=false
    for _, v := range m {
        if v < k {
            invalidFound = true
        }
    }
    if !invalidFound {
        return len(s)
    }

    start :=0
    maxlen:=0
    for i, e := range s {
        if m[e] < k {
            if i > start {
                maxlen=max(maxlen, longestSubstring(s[start:i], k))
            }
            start = i+1
        }
    }

    if start < len(s) {
        maxlen=max(maxlen, longestSubstring(s[start:], k))
    }

    return maxlen
}