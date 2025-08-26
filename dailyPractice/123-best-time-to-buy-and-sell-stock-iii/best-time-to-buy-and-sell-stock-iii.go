func maxProfit(prices []int) int {
    k:=2
    prev, curr := make([]int, 2*k), make([]int, 2*k)
    for t := 0; t < 2*k; t++ {
		if t%2 == 0 { // buy state
			prev[t] = -prices[0]
		}
	}

    for i := 1; i < len(prices); i++ {
        for trans:=0; trans<2*k; trans++ {
            if trans%2 == 0{
                // buy
                if trans ==0 {
                    curr[trans] = max(-prices[i], prev[trans])
                } else {
                curr[trans] = max(-prices[i]+prev[trans-1], prev[trans])
                }
            } else {
                curr[trans] = max(prices[i]+prev[trans-1], prev[trans])
            }
        }
        copy(prev, curr)
    }

    // check all sell states
    res := 0
    for i:=1; i<2*k; i+=2 {
        res = max(res, prev[i])
    }
    return res
}