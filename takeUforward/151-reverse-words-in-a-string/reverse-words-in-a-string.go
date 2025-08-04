func solve(b []byte) []byte {
    i, j := 0, len(b)-1
    for i<j {
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }

    return b
}

func reverseWords(s string) string {
    // remove leading and trailing spaces
    l := 0
    for l < len(s) && s[l] == ' '{
        l++
    }

    s=s[l:]
    r:=len(s)-1 
    for r >= 0 && s[r] == ' '{
        r--
    }
    s=s[:r+1]

    bs := solve([]byte(s))
    start := 0
    for i:=0; i<len(bs); i++ {
        //fmt.Println(i)
        if rune(bs[i]) == ' ' {
            //fmt.Println(start)
            solve(bs[start:i])
            j := i+1
            for j<len(bs) && bs[i] == bs[j] {
                j++
            }
            bs=append(bs[:i+1], bs[j:]...)
            start = i+1
        }
    }

    solve(bs[start:len(bs)])
    return string(bs)
}