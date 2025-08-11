func findSumTill(n int) int {
    if n <= 0 {
        return 0
    }
    return n*(n+1)/2
}

func createPower(n int) ([]int) {
    powers := make([]int, 32, 32)
    leftShift := 32
    curr := 1<<leftShift
    for curr > n {
        leftShift--
        curr = 1<<leftShift
    }

    i:=31
    for n>0 && leftShift>=0 {
        val := 1<<leftShift
        for n-val >= 0 {
            powers[i] = leftShift
            n-=val
            i--
        } 
        leftShift--
    }

    powers=powers[i+1:]
    return powers
}

func modPow(base, exp, mod int) int {
	res := 1
	for exp > 0 {
		if exp&1 != 0 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}

func productQueries(n int, queries [][]int) []int {
    powers := createPower(n)
    fmt.Println(powers)
    for i:=1; i<len(powers); i++ {
        powers[i]+=powers[i-1]
    }
    fmt.Println(powers)
    res := make([]int, len(queries))

    for i, q := range queries {
        l, r := q[0], q[1]
		totalExp := powers[r]
		if l > 0 {
			totalExp -= powers[l-1]
		}
        res[i] = modPow(2, totalExp, 1000000007)
        // res[i] = (1<<totalExp)%1000000007
    }
    return res
}