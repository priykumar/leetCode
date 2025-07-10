func singleNonDuplicate(A []int) int {
    n := len(A)
    if n==1 {
        return A[0]
    }
    
    if A[0] != A[1] {
        return A[0]
    }
    if A[n-1] != A[n-2] {
        return A[n-1]
    }

    l, m, r := 1, 0, n-2

    for l<=r {
        m = l+(r-l)/2
        if A[m] != A[m-1] && A[m] != A[m+1] {
            return A[m]
        }

        if m%2 == 0 {
            if A[m] == A[m-1] {
                r=m-1
            } else {
                l=m+1
            }
        } else {
            if A[m] == A[m-1] {
                l=m+1
            } else {
                r=m-1
            }
        }
    }

    return -1
}