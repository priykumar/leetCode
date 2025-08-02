var dp = map[string]bool{}

func solve(start int, word string, wordMap map[string]bool) bool {
    if start>=len(word) {
        return false
    }

    if val, exist := dp[word[start:]]; exist {
        return val
    }

    if wordMap[word[start:]] == true {
        dp[word[start:]] = true
        return true
    }

    for j:=start; j<len(word); j++ {
        if wordMap[word[start:j+1]] == true {
            if solve(j+1, word, wordMap) {
                dp[word[start:j+1]] = true
                return true
            }
        }
    }

    dp[word[start:]] = false
    return false
}

func wordBreak(s string, wordDict []string) bool {
    wordMap := map[string]bool{}
    dp = make(map[string]bool)
    for _, word := range wordDict {
        wordMap[word]=true
    }

    return solve(0, s, wordMap)
}