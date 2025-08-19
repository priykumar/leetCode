package main

import "fmt"

func RequiredChildFor_m_pages(pages []int, m int) int {
	count := 1
	currPageCount := 0
	for _, p := range pages {
		if currPageCount+p <= m {
			currPageCount += p
		} else {
			count++
			currPageCount = p
		}
	}

	return count
}

func solve(pages []int, child int) int {
	maxPage, totalPage := pages[0], 0
	for _, p := range pages {
		maxPage = max(maxPage, p)
		totalPage += p
	}

	l, r := maxPage, totalPage
	res := -1
	for l <= r {
		m := l + (r-l)/2
		reqChild := RequiredChildFor_m_pages(pages, m)
		if reqChild <= child {
			res = m
			r = m - 1 // try smaller max
		} else {
			l = m + 1 // need more capacity
		}
	}

	return res
}

func main() {
	a := []int{12, 34, 67, 90}
	m := 2

	fmt.Println(solve(a, m))
}
