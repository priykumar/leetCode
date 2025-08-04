var count int64
func merge(a []int, l, m, r int) {
    // Count reverse pairs
	j := m + 1
	for i := l; i <= m; i++ {
		for j <= r && a[i] > 2*a[j] {
			j++
		}
		count += int64(j - (m + 1))
	}


	A := make([]int, m+1-l)
	B := make([]int, r-m)
	copy(A, a[l:m+1])
	copy(B, a[m+1:r+1])

	lenA, lenB := len(A), len(B)

	i, j, k := 0, 0, l
	for i < lenA && j < lenB {
		if A[i] < B[j] {
			a[k] = A[i]
			i++
		} else {
			a[k] = B[j]
			j++
		}
		k++
	}

	for i < lenA {
		a[k] = A[i]
		i++
		k++
	}

	for j < lenB {
		a[k] = B[j]
		j++
		k++
	}
}

func solve(a []int, l, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	solve(a, l, m)
	solve(a, m+1, r)

	merge(a, l, m, r)
}

func reversePairs(nums []int) int {
    count = 0
	solve(nums, 0, len(nums)-1)
    fmt.Println(count)
    return int(count)
}