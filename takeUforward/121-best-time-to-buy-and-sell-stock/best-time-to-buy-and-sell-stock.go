func maxProfit(prices []int) int {
    minPrice:=prices[0]
    maxProfit:=0
    for i:=1; i<len(prices); i++ {
        minPrice = min(minPrice, prices[i])
        maxProfit = max(maxProfit, prices[i]-minPrice)
    }

    return maxProfit
}