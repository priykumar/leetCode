package main

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func solve(s string, start, n int) [][]string {
	if start == n-1 {
		return [][]string{{s[start:]}}
	}

	for i := start; i < n; i++ {
		// parition
		p1 := isPalindrome(s[start:i])
		p2 := isPalindrome(s[i:])

		solve(s, i, n)

	}
}

func main() {
	s := "aabaa"
	solve(s, 0, len(s))
}
