func putMarbles(weights []int, k int) int64 {
    n := len(weights)
    if n == k || k == 1{
        return 0
    }
    pairSum := make([]int, n-1)
    for i:=0; i<n-1; i++ {
        pairSum[i] = weights[i] + weights[i+1]
    }
    sort.Ints(pairSum)
    var minVal, maxVal int64 //:= 0, 0
    for i:=0; i<k-1; i++ {
        minVal=minVal+int64(pairSum[i])
        maxVal=maxVal+int64(pairSum[n-1-1-i])
    } 

    return maxVal-minVal
}