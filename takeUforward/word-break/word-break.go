package main

import "fmt"

var wordMap map[string]bool
var dp map[int]bool

func solve(s string, start int) bool {
	if start >= len(s) {
		return true
	}

	if _, exist := dp[start]; exist {
		return exist
	}

	for i := start; i < len(s); i++ {
		if wordMap[s[start:i+1]] {
			if solve(s, i+1) {
				dp[start] = true
				return true
			}
		}
	}

	dp[start] = false
	return false
}

func populateMap(wordDict []string) {
	dp = make(map[int]bool)
	wordMap = make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}
}

func main() {
	s := "takeuforward"
	wordDict := []string{"take", "forward", "you", "u"}
	populateMap(wordDict)
	fmt.Println(solve(s, 0))

	s = "takeuforward"
	wordDict = []string{"takeu", "orward", "you", "fo"}
	populateMap(wordDict)
	fmt.Println(solve(s, 0))
}
