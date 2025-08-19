package main

import "fmt"

func solve(M [][]int) int {
	personCount := len(M)
	l, r := 0, personCount-1
	for l < r {
		if M[l][r] == 1 { // l knows r, so l can't be a celeb
			l++
		} else { // l doesn't know r, so r can't be a celeb
			r--
		}
	}

	leftPerson := l
	// check if eveyone else know leftPerson
	for i := 0; i < personCount; i++ {
		if l != leftPerson && M[i][leftPerson] == 0 {
			return -1
		}
	}

	// check if leftPerson doesn't know anyone
	for i := 0; i < personCount; i++ {
		if M[leftPerson][i] == 1 {
			return -1
		}
	}

	return leftPerson
}

func main() {
	M := [][]int{{0, 1, 1, 0}, {0, 0, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}}
	fmt.Println(solve(M))
}
