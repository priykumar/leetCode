package main

import (
	"fmt"
	"sort"
)

func solve(Arrival, Departure []int) {
	N := len(Arrival)
	a, d := 0, 0

	sort.Ints(Departure)
	count := 0
	res := 0
	for a < N && d < N {
		if Arrival[a] <= Departure[d] {
			count++
			a++
			res = max(res, count)
		} else {
			count--
			d++
		}
	}
	fmt.Println(res)
}

func main() {
	Arrival := []int{900, 940, 950, 1100, 1500, 1800}
	Departure := []int{910, 1200, 1120, 1130, 1900, 2000}
	solve(Arrival, Departure)

	Arrival = []int{900, 1100, 1235}
	Departure = []int{1000, 1200, 1240}
	solve(Arrival, Departure)
}
