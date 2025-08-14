func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // sort to handle duplicates
	var res [][]int
	used := make([]bool, len(nums))
	var backtrack func(path []int)

	backtrack = func(path []int) {
		if len(path) == len(nums) {
			// make a copy before adding
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// Skip duplicates
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtrack(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]int{})
	return res
}