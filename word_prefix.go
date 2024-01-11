package main

import "fmt"

type Trinode struct {
	children map[rune]*Trinode
	isEnd    bool
}

type Trie struct {
	root *Trinode
}

func NewTrie() *Trie {
	return &Trie{root: &Trinode{
		children: make(map[rune]*Trinode),
	}}
}

func (t *Trie) Insert(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &Trinode{
				children: make(map[rune]*Trinode)}

		}
		node = node.children[char]
		if node.isEnd {
			return true //prefix found
		}
	}
	if len(node.children) > 0 {
		return true
	}
	node.isEnd = true
	return false

}

func noPrefix(words []string) {
	Trie := NewTrie()

	for _, word := range words {
		if Trie.Insert(word) {
			fmt.Println("BAD SET\n")
			fmt.Println(word)
			return

		}
	}
	fmt.Println("GOOD SET")
}

func main() {
	// Example usage:
	words1 := []string{"aabcde", "aab", "defgab", "abcde", "cedaaa", "bbbbbbbbbb", "jabjjjad"}
	noPrefix(words1)

	words2 := []string{"aab", "aac", "aacghgh", "aabghgh"}
	noPrefix(words2)
}
