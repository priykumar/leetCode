type Trie struct {
    children [2]*Trie
}

func NewTrie() Trie {
    return Trie {
        children: [2]*Trie{},
    }
}

func (t *Trie)Insert(ele int) {
    curr := t
    for i:=31; i>=0; i-- {
        bit := (ele>>i)&1
        if curr.children[bit] == nil {
            n := NewTrie()
            curr.children[bit] = &n
        }
        curr=curr.children[bit]
    }
}

func (t *Trie)Traverse(ele int) int{
    curr:=t
    res:=0
    for i:=31; i>=0; i-- {
        bit := (ele >> i)&1
        // If opposite bit exist then XOR with that opposite bit gives 1, so add 1 to the result to left shift
        if curr.children[bit^1] != nil {
            curr=curr.children[bit^1]
            res |= (1 << i)
        } else {
            curr=curr.children[bit]
        }  
    }

    return res
}

func maximizeXor(nums []int, queries [][]int) []int {
    for i:=0; i<len(queries); i++ {
        queries[i] = append(queries[i], i)
    }

    // sort according to mi
    sort.SliceStable(queries, func(a, b int) bool{
        return queries[a][1] < queries[b][1]
    })

    // sort nums
    sort.Ints(nums)

    t:=NewTrie()

    currNums:=0
    lenNums:=len(nums)
    res:=make(map[int]int)
    for i:=0; i<len(queries); i++ {
        for currNums < lenNums && nums[currNums] <= queries[i][1] {
            t.Insert(nums[currNums])
            currNums++
        }

        if currNums>0 {
            res[queries[i][2]] = t.Traverse(queries[i][0])
        } else {
            res[queries[i][2]]=-1
        }
    }

    l:=len(queries)
    result:=[]int{}
    for i:=0; i<l; i++ {
        result=append(result, res[i])
    }

    return result
}