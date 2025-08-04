func totalFruit(fruits []int) int {
    if len(fruits) <= 2 {
        return len(fruits)
    }

    EleLastIdx := make(map[int]int)
    UniqueEle := make(map[int]bool)
    count := make(map[int]int)
    EleLastIdx[fruits[0]] = 0
    EleLastIdx[fruits[1]] = 1
    UniqueEle[fruits[0]]=true
    UniqueEle[fruits[1]]=true
    count[fruits[0]]+=1
    count[fruits[1]]+=1

    res := 0
    for i:=2; i<len(fruits); i++ {
        // If kind of fruit already exist in bucket
        if UniqueEle[fruits[i]] == true {
            EleLastIdx[fruits[i]] = i
            count[fruits[i]]++
        } else if len(UniqueEle) < 2 {
            // If bucket has less than 2 kind of fruits
            UniqueEle[fruits[i]] = true
            EleLastIdx[fruits[i]] = i
            count[fruits[i]]++
        } else {
            prevEle := fruits[i-1]
            var eleToBeRemoved int
            for k, _ := range count {
                if k != prevEle {
                    eleToBeRemoved = k
                    break
                }
            }
            fmt.Println("Removing ", eleToBeRemoved, res, count)
            res = max(res, count[eleToBeRemoved] + count[prevEle])

            count[prevEle]=(i-1)-EleLastIdx[eleToBeRemoved]
            EleLastIdx[fruits[i]] = i
            // remove eleToBeRemoved
            delete(count, eleToBeRemoved)
            delete(UniqueEle, eleToBeRemoved)

            UniqueEle[fruits[i]] = true
            count[fruits[i]]++
        }
    }

    tempRes := 0
    for _, v := range count {
        tempRes += v
    }

    res=max(res, tempRes)

    return res
}