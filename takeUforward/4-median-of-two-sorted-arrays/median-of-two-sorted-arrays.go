func findMedianSortedArrays(a []int, b []int) float64 {
    n1, n2 := len(a),len(b)
    if n1 > n2 {
        return findMedianSortedArrays(b, a)   // take first array as the smaller one
    }
    low, high := 0, n1
    left := (n1+n2+1)/2 // how many elements should be on the left side of the partition
    for low <= high {
        mid1:=(low+high)/2
        mid2:=left-mid1
        l1, l2, r1, r2 := -1<<30, -1<<30, 1<<31, 1<<31
        if mid1 < n1 {
            r1 = a[mid1]
        }
        if mid2 < n2 {
            r2 = b[mid2]
        }
        if mid1-1>=0 {
            l1 = a[mid1-1]
        }
        if mid2-1>=0 {
            l2 = b[mid2-1]
        }

        if l1<=r2 && l2<=r1 {
            if (n1+n2)%2 == 0 {
                return float64(max(l1,l2)+min(r1,r2))/2.0
            } else {
                return float64(max(l1,l2))
            }
        }

        if l1 > r2 {
            high=mid1-1
        } else {
            low = mid1+1
        }
    }

    return -1.0
}