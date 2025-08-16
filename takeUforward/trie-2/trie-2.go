// Implement Trie - II
// Problem Statement: Implement a Trie data structure that supports the following methods:
// Insert (word): To insert a string `word` in the Trie.
// Count Words Equal To (word): Return the count of occurrences of the string word in the Trie.
// Count Words Starting With (prefix): Return the count of words in the Trie that have the string “prefix” as a prefix.
// Erase (word): Delete one occurrence of the string word from the Trie.

package main

import "fmt"

type Trie struct {
	prefixCount int
	wordCount   int
	children    []*Trie
}

func Constructor() Trie {
	return Trie{
		prefixCount: 0,
		wordCount:   0,
		children:    make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	fmt.Println("Insert ", word)
	this.prefixCount++
	for _, ele := range word {
		if this.children[ele-'a'] == nil {
			newNode := Constructor()
			this.children[ele-'a'] = &newNode
		}
		this = this.children[ele-'a']
		this.prefixCount++
	}
	this.wordCount++
}

func (this *Trie) CountWord(word string) int {
	fmt.Println("Count word ", word)
	for _, ele := range word {
		if this.children[ele-'a'] == nil {
			return 0
		}
		this = this.children[ele-'a']
	}
	return this.wordCount
}

func (this *Trie) CountPrefix(word string) int {
	fmt.Println("Count prefix ", word)
	for _, ele := range word {
		if this.children[ele-'a'] == nil {
			return 0
		}
		this = this.children[ele-'a']
	}
	return this.prefixCount
}

func (this *Trie) DeleteWord(word string) {
	fmt.Println("Delete ", word)
	for _, ele := range word {
		if this.prefixCount <= 1 {
			this.children[ele-'a'] = nil
			return
		}
		this.prefixCount--
		this = this.children[ele-'a']
	}

	if this.wordCount >= 1 {
		this.wordCount--
	}
}

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println("Expected: 1\tGot: ", trie.CountWord("apple")) // 1
	fmt.Println("Expected: 1\tGot: ", trie.CountPrefix("app")) // 1

	trie.Insert("apple")
	trie.Insert("apple")
	fmt.Println("Expected: 3\tGot: ", trie.CountWord("apple")) // 3
	fmt.Println("Expected: 3\tGot: ", trie.CountPrefix("app")) // 3

	trie.Insert("app")
	fmt.Println("Expected: 1\tGot: ", trie.CountWord("app"))   // 1
	fmt.Println("Expected: 4\tGot: ", trie.CountPrefix("app")) // 4

	trie.DeleteWord("apple")
	fmt.Println("Expected: 2\tGot: ", trie.CountWord("apple")) // 2
	fmt.Println("Expected: 3\tGot: ", trie.CountPrefix("app")) // 3

	trie.DeleteWord("apple")
	trie.DeleteWord("apple")
	fmt.Println("Expected: 0\tGot: ", trie.CountWord("apple")) // 0
	fmt.Println("Expected: 1\tGot: ", trie.CountPrefix("app")) // 1

	trie.DeleteWord("banana")
	fmt.Println("Expected: 0\tGot: ", trie.CountWord("banana")) // 0

	trie.Insert("apple")
	trie.Insert("apricot")
	fmt.Println("Expected: 3\tGot: ", trie.CountPrefix("ap")) // 3
	trie.DeleteWord("apple")
	fmt.Println("Expected: 0\tGot: ", trie.CountWord("apple"))   // 0
	fmt.Println("Expected: 1\tGot: ", trie.CountWord("apricot")) // 1
	fmt.Println("Expected: 2\tGot: ", trie.CountPrefix("ap"))    // 2

	trie.DeleteWord("app")
	fmt.Println("Expected: 0\tGot: ", trie.CountPrefix("app")) // 0
}
