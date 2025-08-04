package main

import "fmt"

var count int

func merge(a []int, l, m, r int) {
	A := make([]int, m+1-l)
	B := make([]int, r-m)
	copy(A, a[l:m+1])
	copy(B, a[m+1:r+1])

	lenA, lenB := len(A), len(B)

	i, j, k := 0, 0, l
	for i < lenA && j < lenB {
		if A[i] <= B[j] {
			a[k] = A[i]
			i++
		} else {
			a[k] = B[j]
			count += lenA - i
			j++
		}
		k++
	}

	for i < lenA {
		a[k] = A[i]
		i++
		k++
	}

	for j < lenB {
		a[k] = B[j]
		j++
		k++
	}
}

func solve(a []int, l, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	solve(a, l, m)
	solve(a, m+1, r)

	merge(a, l, m, r)
}

func main() {
	a := []int{2, 3, 7, 1, 3, 5}
	count = 0
	solve(a, 0, len(a)-1)
	fmt.Println(count)

	a = []int{-10, -5, 6, 11, 15, 17}
	count = 0
	solve(a, 0, len(a)-1)
	fmt.Println(count)

	a = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	count = 0
	solve(a, 0, len(a)-1)
	fmt.Println(count)
}
