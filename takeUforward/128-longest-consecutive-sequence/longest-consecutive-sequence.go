func longestConsecutive(A []int) int {
    m:=map[int]int{}
    for _, v:=range A{
        m[v]++
    }
    
    result:=0
    for key, _ := range m {
        if m[key-1] == 0 {
            tempLen:=0
            startEle:=key
            for m[startEle]>0 {
                tempLen++ 
                startEle++
            }
            if tempLen > result {
                result=tempLen
            }
        }
    }
    
    return result
}