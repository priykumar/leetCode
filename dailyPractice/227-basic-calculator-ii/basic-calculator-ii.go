func calculate(s string) int {
    // Precedence of / and * are same > Precedence of + and - are same
    curr := 0
    stack := []int{}
    lastSign := '+'
    for i, ele := range s {
        if ele>='0' && ele<='9' {
            curr = (curr*10)+int(ele-'0')
        }
        if ele=='+' || ele=='-' || ele=='*'|| ele=='/' || i == len(s)-1 {
            if lastSign == '+' {
                stack=append(stack, curr)
            } else if lastSign == '-' {
                stack=append(stack, -curr)  
            } else if lastSign == '*' {
                lastEle := stack[len(stack)-1]
                stack[len(stack)-1] = lastEle*curr
            } else if lastSign == '/' {
                lastEle := stack[len(stack)-1]
                stack[len(stack)-1] = lastEle/curr
            }
            lastSign = ele
            curr=0
        }
    }

    res:=0
    for i:=0; i<len(stack); i++ {
        res+=stack[i]
    }

    return res
}