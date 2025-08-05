func numOfUnplacedFruits(fruits []int, baskets []int) int {
    for _, f := range fruits {
        for j, b := range baskets {
            if b >= f {
                baskets[j]=-1
                // fmt.Println(j)
                break
            }
        }
    }

    //fmt.Println(baskets)
    count:=0
    for i:=0; i<len(baskets); i++ {
        if baskets[i]!=-1 {
            count++
        }
    }
    return count
}