func longestPalindrome(s string) string {
    if len(s) <=1 {
        return s
    }

    maxlen, res := 0, ""
    for i:=0; i<len(s); i++ {
        l1, r1 := expandFromCenter(s, i, i)
        l2, r2 := expandFromCenter(s, i, i+1)

        if r1-l1+1 > maxlen {
            maxlen = r1-l1+1
            res = s[l1:r1+1]
        }

        if r2-l2+1 > maxlen {
            maxlen = r2-l2+1
            res = s[l2:r2+1]
        }
    }

    return res
}

func expandFromCenter(s string, i, j int) (int, int) {
    for i>=0 && j<len(s) && s[i] == s[j] {
        i--
        j++
    }

    return i+1, j-1
}