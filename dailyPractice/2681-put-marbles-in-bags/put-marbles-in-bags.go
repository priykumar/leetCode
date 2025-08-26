// TC: O(logn)
func putMarbles(weights []int, k int) int64 {
    // https://www.youtube.com/watch?v=RyJpH8cghrE
    n := len(weights)
    if n == k || k == 1{
        return 0
    }
    pairSum := make([]int, n-1)
    for i:=0; i<n-1; i++ {
        pairSum[i] = weights[i] + weights[i+1]
    }
    sort.Ints(pairSum) // If we use minheap and maxheap to get k-1 largest and k-1 smallest, then TC would be little less (O(klogn))
    var minVal, maxVal int64
    for i:=0; i<k-1; i++ {
        minVal=minVal+int64(pairSum[i])
        maxVal=maxVal+int64(pairSum[n-1-1-i])
    } 

    return maxVal-minVal
}