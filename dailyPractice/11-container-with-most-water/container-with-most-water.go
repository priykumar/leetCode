func maxArea(height []int) int {
    l, r := 0, len(height)-1
    water:=0
    for l < r {
        if height[l] < height[r] {
            water=max(water, (r-l)*height[l])
            l++
        } else if height[l] > height[r] {
            water=max(water, (r-l)*height[r])
            r--
        } else {
            water=max(water, (r-l)*height[r])
            l++
            r--
        }
    }

    return water
}