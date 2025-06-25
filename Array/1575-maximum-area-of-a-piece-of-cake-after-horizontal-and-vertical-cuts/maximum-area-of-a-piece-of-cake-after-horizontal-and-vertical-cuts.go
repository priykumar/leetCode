func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    horizontalCuts = append(horizontalCuts, 0)
    horizontalCuts = append(horizontalCuts, h)
    sort.Ints(horizontalCuts)

    verticalCuts = append(verticalCuts, 0)
    verticalCuts = append(verticalCuts, w)
    sort.Ints(verticalCuts)

    maxDistantHorizontalCut := -1
    for i:=1; i<len(horizontalCuts); i++ {
        maxDistantHorizontalCut=max(maxDistantHorizontalCut, horizontalCuts[i]-horizontalCuts[i-1])
    }

    maxDistantVerticalCut := -1
    for i:=1; i<len(verticalCuts); i++ {
        maxDistantVerticalCut=max(maxDistantVerticalCut, verticalCuts[i]-verticalCuts[i-1])
    }

    return ((maxDistantHorizontalCut%1000000007) * (maxDistantVerticalCut%1000000007))%1000000007
}