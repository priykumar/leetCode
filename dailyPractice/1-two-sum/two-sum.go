func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)
    for i, n := range nums {
        if val, exist := m[target-n]; exist {
            return []int{val, i}
        }
        m[n]=i
    }
    return []int{-1, -1}
}