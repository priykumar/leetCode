package main

import "fmt"

var val2count map[int]int

func solve(nums []int, k int) []int {
	for i := 0; i < len(nums) && i < k; i++ {
		val2count[nums[i]]++
	}

	res := []int{len(val2count)}
	if k >= len(nums) {
		return res
	}

	l, r := 0, k-1
	uniqueCount := len(val2count)
	for r+1 < len(nums) {
		oldEle := nums[l]
		val2count[oldEle]--
		if val2count[oldEle] == 0 {
			delete(val2count, oldEle)
			uniqueCount--
		}
		l++

		r++
		newEle := nums[r]
		if _, exist := val2count[newEle]; !exist {
			uniqueCount++
			val2count[newEle] = 1
		} else {
			val2count[newEle]++
		}
		fmt.Println(val2count)
		res = append(res, uniqueCount)
	}

	return res
}

func main() {
	val2count = make(map[int]int)
	nums := []int{1, 2, 3, 2, 2, 1, 3}
	k := 3
	fmt.Println(solve(nums, k))

	val2count = make(map[int]int)
	nums = []int{1, 1, 1, 1, 2, 3, 4}
	k = 4
	fmt.Println(solve(nums, k))
}
