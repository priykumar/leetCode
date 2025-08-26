func maxCount(banned []int, n int, maxSum int) int {
    bannedSet := make(map[int]bool)
    for _, b := range banned {
        bannedSet[b] = true
    }

    sum, count := 0, 0
    for i := 1; i <= n; i++ {
        if bannedSet[i] {
            continue
        }
        if sum+i > maxSum {
            break
        }
        sum += i
        count++
    }
    return count
}
