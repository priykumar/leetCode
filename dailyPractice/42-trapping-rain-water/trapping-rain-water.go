func trap(height []int) int {
    n:=len(height)
    rmax := make([]int, n)
    rmax[n-1]=height[n-1]
    for i:=n-2; i>=0; i-- {
        rmax[i]=max(rmax[i+1], height[i])
    }

    lmax:=height[0]
    water:=0
    for i:=1; i<n; i++ {
        lmax=max(lmax, height[i])
        tempWater := min(lmax, rmax[i])-height[i]
        water+=max(0, tempWater)
    }

    return water
}