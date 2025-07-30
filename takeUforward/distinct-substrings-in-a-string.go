package main

import "fmt"

type Trie struct {
	children []*Trie
}

func Constructor() Trie {
	return Trie{
		children: make([]*Trie, 26),
	}
}

func (t *Trie) Insert(ele string) int {
	count := 0
	for _, ch := range ele {
		if t.children[ch-'a'] == nil {
			newNode := Constructor()
			t.children[ch-'a'] = &newNode
			count++
		}
		t = t.children[ch-'a']
	}

	return count
}

func main() {
	strings := []string{"aba", "abc", "ab", "a", "abca", "abacaba", "abacabadabacaba"}
	res := 0
	for _, s := range strings {
		res = 0
		trie := Constructor()
		for i := 0; i < len(s); i++ {
			for j := i; j < len(s); j++ {
				tempCount := trie.Insert(s[i : j+1])
				res += tempCount
			}
		}
		fmt.Println(s, " : ", res+1) // 1 FOR THE EMPTY SUBSTRING
	}

}
