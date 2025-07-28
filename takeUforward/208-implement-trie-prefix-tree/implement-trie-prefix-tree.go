type Trie struct {
    isEnd bool
    children []*Trie
}


func Constructor() Trie {
    return Trie {
        isEnd: false,
        children: make([]*Trie, 26),
    }
}

func (this *Trie) Insert(word string)  {
    for _, ele := range word {
        if this.children[ele-'a'] == nil {
            t := Constructor()
            this.children[ele-'a'] = &t
        } 
        this = this.children[ele-'a']
    }
    this.isEnd=true
}

func (this *Trie) Search(word string) bool {
    for _, ele := range word {
        if this.children[ele-'a'] == nil {
            return false
        } 
        this = this.children[ele-'a']
    }
    return this.isEnd == true
}


func (this *Trie) StartsWith(prefix string) bool {
    for _, ele := range prefix {
        if this.children[ele-'a'] == nil {
            return false
        } 
        this = this.children[ele-'a']
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */