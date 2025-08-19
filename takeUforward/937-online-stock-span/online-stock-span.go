type StockSpanner struct {
    prices []int
    stack []int
}


func Constructor() StockSpanner {
    return StockSpanner {
        prices: make([]int, 0),
        stack: make([]int, 0),
    }
}


func (this *StockSpanner) Next(price int) int {
    this.prices = append(this.prices, price)
    for len(this.stack)>0 && this.prices[this.stack[len(this.stack)-1]] <= price {
        this.stack = this.stack[:len(this.stack)-1]
    }

    res := 1
    if len(this.stack) == 0{
        res = len(this.prices)
    } else {
        res = (len(this.prices)-1)-this.stack[len(this.stack)-1]
    }
    this.stack = append(this.stack, len(this.prices)-1)

    return res
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */