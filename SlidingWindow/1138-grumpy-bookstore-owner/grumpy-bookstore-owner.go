func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    oneCount, totalCustomers, happyCustomers := 0, 0, 0
    for i:=0; i<len(customers); i++ {
        totalCustomers+=customers[i]
        if grumpy[i] == 0 {
            happyCustomers+=customers[i]
        } else {
            oneCount++
        }
    }

    if len(customers) <= minutes {
        return totalCustomers
    }

    additionalHappyCustomer := 0
    for i:=0; i<minutes; i++ {
        if grumpy[i] == 1 {
            additionalHappyCustomer+= customers[i]
        }
    }

    temp_additionalHappyCustomer := additionalHappyCustomer
    l, r, n := 0, minutes-1, len(customers)
    for r < n {
        if grumpy[l] == 1 {
            temp_additionalHappyCustomer-=customers[l]
        }
        l++
        r++
        if r<n && grumpy[r] == 1 {
            temp_additionalHappyCustomer+=customers[r]
        }

        additionalHappyCustomer=max(additionalHappyCustomer, temp_additionalHappyCustomer)
    }

    return happyCustomers + additionalHappyCustomer
}