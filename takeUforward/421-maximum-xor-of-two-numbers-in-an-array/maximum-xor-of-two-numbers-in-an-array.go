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

func findMaximumXOR(nums []int) int {
    t:=NewTrie()
    for _, ele := range nums {
        t.Insert(ele)
    }

    res:=0
    for _, ele := range nums {
        res=max(res, t.Traverse(ele))
    }
    return res
}