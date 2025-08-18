package main

import (
	"fmt"
	"sort"
)

type S struct {
	s   int
	e   int
	idx int
}

func main() {
	N := 6
	start := []int{1, 3, 0, 5, 8, 5}
	end := []int{2, 4, 5, 7, 9, 9}

	meetings := make([]S, 0, N)
	for i := 0; i < N; i++ {
		meetings = append(meetings, S{start[i], end[i], i})
	}

	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i].e < meetings[j].e
	})

	count := 1
	res := []int{0}
	for i := 1; i < N; i++ {
		if meetings[i].s >= meetings[i-1].e {
			count++
			res = append(res, meetings[i].idx)
		}
	}

	fmt.Println("Total meetings:", count)
	for i := 0; i < len(res); i++ {
		fmt.Printf("Start: %d\t End: %d\n", start[res[i]], end[res[i]])
	}
}
