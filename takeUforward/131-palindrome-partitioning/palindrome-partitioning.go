var result [][]string
func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i<j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func solve(s string, temp []string) {
    if len(s) == 0 {
        temp1 := make([]string, len(temp))
        copy(temp1, temp)
        result = append(result, temp1)
        return  
    }

    for i:=1; i<=len(s); i++ {
        substring:=s[:i]
        if isPalindrome(substring) {
            temp = append(temp, substring)
            solve(s[i:], temp)
            temp = temp[:len(temp)-1]
        }
    }
}

func partition(s string) [][]string {
    result=[][]string{}
    temp:=[]string{}
    solve(s, temp)
    return result
}