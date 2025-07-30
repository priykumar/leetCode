package main

import "fmt"

func solve(nums []int) int {
	n := len(nums)
	sumArr := []int{nums[0]}
	res := nums[0]
	for i := 1; i < n; i++ {
		tempMaxSum := nums[i]
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && sumArr[j]+nums[i] > tempMaxSum {
				tempMaxSum = sumArr[j] + nums[i]
			}
		}
		sumArr = append(sumArr, tempMaxSum)
		res = max(res, tempMaxSum)
	}
	return res
}

func main() {
	nums := []int{1, 101, 2, 3, 100}
	fmt.Println(nums, "  :  ", solve(nums))

	nums = []int{4, 1, 2, 3}
	fmt.Println(nums, "  :  ", solve(nums))
}
