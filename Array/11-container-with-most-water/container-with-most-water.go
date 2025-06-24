func maxArea(height []int) int {
    maxWater:=-1
    i, j := 0, len(height)-1

    for i<j {
        maxWater=max(maxWater, min(height[i], height[j])*(j-i))

        if height[i] < height[j] {
            i++
        } else if height[i] > height[j] {
            j--
        } else {
            i++
            j--
        }
    }

    return maxWater
}