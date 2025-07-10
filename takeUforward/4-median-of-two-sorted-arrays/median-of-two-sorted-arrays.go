func findMedianSortedArrays(A []int, B []int) float64 {
    l1, l2 := len(A), len(B)

    isEven := false
    midIndex1 := (l1+l2)/2 + 1
    midIndex2 := -1
    if (l1+l2)%2 == 0 {
        isEven = true
        midIndex1--
        midIndex2=midIndex1+1
    }

    midEleSum:=0

    a, b, ele := 0, 0, 0
    count := 0
    for a < l1 && b < l2 {
        if A[a] <= B[b] {
            ele = A[a]
            a++
        } else {
            ele = B[b]
            b++
        }
        count++
        if midIndex1==count {
            midEleSum=ele
            if midIndex2 == -1 {
                break
            }
        }
        if midIndex2==count {
            midEleSum+=ele
            break
        }
    } 

    fmt.Println(midIndex1, midIndex2, midEleSum, count, isEven)
    if count<midIndex1 || (midIndex2!=-1 && count<midIndex2) {
        for a < l1 {
            ele = A[a]
            a++
            count++
        
            if midIndex1==count {
                midEleSum=ele
                if midIndex2 == -1 {
                    break
                }
            }
            if midIndex2==count {
                midEleSum+=ele
                break
            }
        }

        fmt.Println(b)
        for b < l2 {
            ele = B[b]
            b++
            count++

            fmt.Println(b, count)
            if midIndex1==count {
                midEleSum=ele
                if midIndex2 == -1 {
                    fmt.Println("here")
                    break
                }
            }
            if midIndex2==count {
                midEleSum+=ele
                break
            }
        }
    }

    fmt.Println("->", midIndex1, midIndex2, midEleSum, count)
    res := float64(midEleSum)
    if isEven {
        res=res/2
    }

    return res
}