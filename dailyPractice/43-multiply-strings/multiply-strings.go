func multiply(num1 string, num2 string) string {
    if num1=="0" || num2=="0" {
        return "0"
    }
    if len(num2)>len(num1) {
        num1, num2 = num2, num1
    }
    n1, n2 := len(num1), len(num2)

    res := make([]int, n1+n2)
    start:=0
    for j:=n2-1; j>=0; j-- {
        currStart:=n1+n2-1-start
        for i:=n1-1; i>=0; i-- {
            tempRes := int(num2[j]-'0')*int(num1[i]-'0')
            res[currStart]+=tempRes
            currStart--
        }
        start++
    }

    carry:=0
    out:=""
    for i:=n1+n2-1; i>=0; i-- {
        tempRes := (res[i]+carry)%10
        carry = (res[i]+carry)/10
        res[i] = tempRes
        out=fmt.Sprintf("%d", tempRes)+out
    }

    i:=0
    for i<len(out) {
        if string(out[i]) == "0" {
            i++
        } else {
            break
        }
    }
    return out[i:]
}