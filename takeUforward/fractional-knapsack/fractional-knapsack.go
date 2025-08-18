package main

import (
	"fmt"
	"sort"
)

type Item struct {
	val   float64
	wt    float64
	ratio float64
}

func FractionalKnapsack(val []float64, wt []float64, capacity float64) float64 {
	n := len(val)
	items := make([]Item, n)
	for i := 0; i < len(val); i++ {
		items[i] = Item{val[i], wt[i], val[i] / wt[i]}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	i := 0
	var res float64 = 0
	for capacity > 0 && i < len(val) {
		if capacity >= items[i].wt {
			res += items[i].val
			capacity -= items[i].wt
			i++
		} else {
			res += items[i].ratio * capacity
			capacity = 0
		}
	}

	return res
}

func main() {
	val := []float64{60, 100, 120}
	wt := []float64{10, 20, 30}
	capacity := 50
	fmt.Println(FractionalKnapsack(val, wt, float64(capacity)))

	val = []float64{60, 100}
	wt = []float64{10, 20}
	capacity = 50
	fmt.Println(FractionalKnapsack(val, wt, float64(capacity)))

	val = []float64{3, 8, 10, 2, 5}
	wt = []float64{10, 4, 20, 8, 15}
	capacity = 40
	fmt.Println(FractionalKnapsack(val, wt, float64(capacity)))

	val = []float64{10, 20, 30}
	wt = []float64{5, 10, 15}
	capacity = 100
	fmt.Println(FractionalKnapsack(val, wt, float64(capacity)))
}
