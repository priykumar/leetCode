const (
    INT_MAX = 1<<31 - 1
    INT_MIN = -1 << 31
)

func removeleadingSpaces(s *string) {
    n:=len(*s)
    i:=0

    for i < n && (*s)[i]==' '{
        i++
    }
    *s=(*s)[i:]
}

func myAtoi(s string) int {
    removeleadingSpaces(&s)
    if s=="" {
        return 0
    }
    curr:=0
    isneg:=false
    if s[curr]=='-' || s[curr]=='+' {
        if s[curr]=='-' {
            isneg = true
        }
        curr++
    }

    n:=len(s)
    for curr < n && s[curr]=='0'{
        curr++
    }

    s=s[curr:]
    curr=0

    var res int64 = 0 
    for curr < len(s) && s[curr]>='0' && s[curr]<='9' {
        num := int64(s[curr]-'0')
        res = (res*10)+num
        curr++

        if !isneg && res > INT_MAX {
            return INT_MAX
        }
        if isneg && -res < INT_MIN {
            return INT_MIN
        }
    }

    if isneg {
        res*=-1
    }

    if !isneg && res > INT_MAX {
        return INT_MAX
    }
    if isneg && res < INT_MIN {
        return INT_MIN
    }
    return int(res)
}