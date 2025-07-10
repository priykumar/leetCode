func trap(height []int) int {
    n := len(height)
    rMax := make([]int, n)
    rMax[n-1] = height[n-1]

    for i:=n-2; i>=0; i-- {
        rMax[i] = max(rMax[i+1], height[i])
    }

    lMax := height[0]
    waterStored := 0
    for i:=1; i<n; i++ {
        lMax = max(lMax, height[i])
        heightDiff := max(0, min(lMax, rMax[i]) - height[i])
        waterStored += heightDiff
    }

    return waterStored
}