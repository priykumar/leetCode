type Bank struct {
    accBalance []int64
    count int
}


func Constructor(balance []int64) Bank {
    b := Bank{}
    b.accBalance = make([]int64, len(balance))
    b.count = len(balance)
    copy(b.accBalance, balance)
    return b
}


func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
    if account1>this.count || account2>this.count {
        return false
    }
    if this.accBalance[account1-1] >= money {
        this.accBalance[account1-1]-=money
        this.accBalance[account2-1]+=money
        return true
    }

    return false
}


func (this *Bank) Deposit(account int, money int64) bool {
    if account>this.count {
        return false
    }
    this.accBalance[account-1]+=money
    return true
}


func (this *Bank) Withdraw(account int, money int64) bool {
    if account>this.count {
        return false
    }
    if this.accBalance[account-1] >= money {
        this.accBalance[account-1]-=money
        return true
    }
    return false
}


/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */