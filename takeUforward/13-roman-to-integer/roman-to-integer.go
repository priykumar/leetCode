func romanToInt(s string) int {
    m:=map[rune]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }

    res:=0

    for i:=0; i<len(s); i++ {
        if i+1<len(s) && m[rune(s[i])]<m[rune(s[i+1])] {
            res+=(m[rune(s[i+1])]-m[rune(s[i])])
            i++
            continue
        }
        res+=m[rune(s[i])]
    }

    return res
}