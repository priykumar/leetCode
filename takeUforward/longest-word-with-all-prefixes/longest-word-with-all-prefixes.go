// Given a string array nums of length n. A string is called a complete string if every prefix of this string is also present in the array nums. Find the longest complete string in the array nums.
// If there are multiple strings with the same length, return the lexicographically smallest one and if no string exists, return "None" (without quotes).
// Input : nums = [ "n", "ni", "nin", "ninj" , "ninja" , "nil" ]
// Output : ninja

// Explanation : The word "ninja" is the longest word which has all its prefixes present in the array.

package main

import "fmt"

type Trie1 struct {
	isEnd    bool
	children []*Trie1
}

func Constructor1() Trie1 {
	return Trie1{
		isEnd:    false,
		children: make([]*Trie1, 26),
	}
}

func (this *Trie1) Insert(word string) {
	fmt.Println("Insert ", word)
	for _, ele := range word {
		if this.children[ele-'a'] == nil {
			newNode := Constructor1()
			this.children[ele-'a'] = &newNode
		}
		this = this.children[ele-'a']
	}
	this.isEnd = true
}

func (this *Trie1) CheckIfAllPrefixPresent(word string) bool {
	for _, ele := range word {
		if this.children[ele-'a'].isEnd == false {
			return false
		}

		this = this.children[ele-'a']
	}

	return true
}

func (this *Trie1) Lookup(arr []string) string {
	maxlen := 0
	res := ""

	for _, ele := range arr {
		if this.CheckIfAllPrefixPresent(ele) {
			if maxlen < len(ele) {
				res = ele
				maxlen = len(ele)
			}
		}
	}

	return res
}

func main() {
	nums := []string{"n", "ni", "nin", "ninj", "ninja", "nil"}

	trie := Constructor1()
	for _, ele := range nums {
		trie.Insert(ele)
	}

	fmt.Println("Result: ", trie.Lookup(nums))
	fmt.Println()

	nums = []string{"ninja", "night", "nil"}
	trie1 := Constructor1()
	for _, ele := range nums {
		trie1.Insert(ele)
	}
	fmt.Println("Result1: ", trie1.Lookup(nums))
}
