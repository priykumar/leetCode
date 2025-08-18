package main

import "fmt"

var res []string
var dx []int = []int{1, 0, -1, 0}
var dy []int = []int{0, 1, 0, -1}
var idx2dir map[int]string = map[int]string{
	0: "D",
	1: "R",
	2: "U",
	3: "L",
}

func solve(grid [][]int, n int, ux, uy int, curr string) {
	if ux == n-1 && uy == n-1 {
		res = append(res, curr)
		return
	}

	for idx := 0; idx < 4; idx++ {
		vx, vy := ux+dx[idx], uy+dy[idx]
		if vx >= 0 && vx < n && vy >= 0 && vy < n && grid[vx][vy] == 1 {
			origGridVal := grid[vx][vy]
			grid[vx][vy] = 2
			curr = curr + idx2dir[idx]
			solve(grid, n, vx, vy, curr)
			curr = curr[:len(curr)-1]
			grid[vx][vy] = origGridVal
		}
	}
}

func main() {
	n := 4
	grid := [][]int{{1, 0, 0, 0}, {1, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 1, 1}}
	res = []string{}
	solve(grid, n, 0, 0, "")
	fmt.Println(res)

	n = 2
	grid = [][]int{{1, 0}, {1, 0}}
	res = []string{}
	solve(grid, n, 0, 0, "")
	fmt.Println(res)
}
