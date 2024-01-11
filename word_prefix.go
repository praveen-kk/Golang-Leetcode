/*
There is a given list of strings where each string contains only lowercase letters from , inclusive. The set of strings is said to be a GOOD SET if no string is a prefix of another string. In this case, print GOOD SET. Otherwise, print BAD SET on the first line followed by the string being checked.

Note If two strings are identical, they are prefixes of each other.

Example

Here 'abcd' is a prefix of 'abcde' and 'bcd' is a prefix of 'bcde'. Since 'abcde' is tested first, print

BAD SET
abcde
.

No string is a prefix of another so print

GOOD SET
Function Description
Complete the noPrefix function in the editor below.

noPrefix has the following parameter(s):
- string words[n]: an array of strings

Prints
- string(s): either GOOD SET or BAD SET on one line followed by the word on the next line. No return value is expected.

Input Format
First line contains , the size of .
Then next  lines each contain a string, .

Constraints

 the length of words[i]
All letters in  are in the range 'a' through 'j', inclusive.

Sample Input00

STDIN       Function
-----       --------
7            words[] size n = 7
aab          words = ['aab', 'defgab', 'abcde', 'aabcde', 'bbbbbbbbbb', 'jabjjjad']
defgab
abcde
aabcde
cedaaa
bbbbbbbbbb
jabjjjad
Sample Output00

BAD SET
aabcde
Explanation
'aab' is prefix of 'aabcde' so it is a BAD SET and fails at string 'aabcde'.

Sample Input01

4
aab
aac
aacghgh
aabghgh
Sample Output01

BAD SET
aacghgh
Explanation
'aab' is a prefix of 'aabghgh', and aac' is prefix of 'aacghgh'. The set is a BAD SET. 'aacghgh' is tested before 'aabghgh', so and it fails at 'aacghgh'.

####################
time complexity - o(s) - where s is the number of characters of the words , in the worst case scenarios it should iterate over all the characters of the words to insert in the trie ds
Space Complexity - o(s) - where s is the number of characters of the words where store them in trie - worst case where no common prefixes exist
*/

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
