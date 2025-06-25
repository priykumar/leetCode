func numPairsDivisibleBy60(time []int) int {
    modArr := make([]int, 60)
    for i:=0; i<len(time); i++ {
        modArr[time[i]%60]++
    }

    i, j := 1, 59
    count:=0
    for i<j {
        count+=modArr[i]*modArr[j]
        i++
        j--
    }
    
    tempCount:=modArr[0]*(modArr[0]-1) + modArr[30]*(modArr[30]-1)
    count+=tempCount/2

    return count
}