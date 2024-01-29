/*
Initial Thought :
Loop though string search in each input line by slicing the line for the max length words of words in number to word map.

Optimized Approach :
create a Trie DS with the words to  1..9 map.

insert the map and iterate through each line and search for the first and last digits.

Time complexity: o(M * N)  - M is the number of lines and N is the average length of each line
Space Complexity: o(1)
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Trinode struct {
	children map[rune]*Trinode
	isEnd    bool
}

type Trie struct {
	root *Trinode
}

func NewTrie() *Trie {
	return &Trie{root: &Trinode{
		children: make(map[rune]*Trinode)}}

}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		//if node.children == nil {
		//	node.children = make(map[rune]*Trinode)
		//}
		if _, exists := node.children[char]; !exists {
			node.children[char] = &Trinode{
				children: make(map[rune]*Trinode)}
		}
		node = node.children[char]
	}
	node.isEnd = true

}
func (t *Trie) digit_calibration(inputLine string) (firstDigit, secondDigit string) {

	firstDigit = t.find_first_digit(inputLine)
	secondDigit = t.findSecondDigit(inputLine)

	return firstDigit, secondDigit

}

func (t *Trie) find_first_digit(input string) string {
	for i := 0; i < len(input); i++ {
		node := t.root
		for j := i; j < len(input); j++ {
			char := rune(input[j])
			if node.children == nil {
				break
			}
			if _, exists := node.children[char]; !exists {
				break
			}
			node = node.children[char]
			if node.isEnd {
				return input[i : j+1]
			}

		}

	}
	return ""

}

func (t *Trie) findSecondDigit(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		node := t.root
		//fmt.Printf("node :  %d\n", i)
		for j := i; j < len(input); j++ { // <-- Corrected index from i to j
			char := rune(input[j])
			if node.children == nil {
				break
			}
			if _, exists := node.children[char]; !exists {
				break
			}
			node = node.children[char]
			if node.isEnd {
				return input[i : j+1]
			}
		}
	}
	return ""
}

func main() {
	trie := NewTrie()
	first := 0
	last := 0
	words := map[string]string{"1": "one", "2": "two", "3": "three", "4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine"}

	// Insert words into the trie, comment words if only integers is to be found
	for key, word := range words {
		trie.Insert(word)
		trie.Insert(key)
	}

	file, err := os.Open("calibration.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	sum := 0

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//var firstDigit, secondDigit string
		line := scanner.Text()

		// Calculate the calibration value for the current line
		firstDigit, secondDigit := trie.digit_calibration(line)
		if len(firstDigit) > 1 {
			for k, v := range words {
				if v == firstDigit {
					first, _ = strconv.Atoi(k)
					break

				}

			}

		} else {
			first, _ = strconv.Atoi(firstDigit)
		}
		if len(secondDigit) > 1 {
			for k, v := range words {
				if v == secondDigit {
					last, _ = strconv.Atoi(k)
					break

				}

			}

		} else {
			last, _ = strconv.Atoi(secondDigit)
		}

		digit := first*10 + last
		fmt.Printf("numbers %d\n", digit)

		// Add the calibration value to the sum
		sum += digit
	}
	fmt.Printf("sum %d", sum)

}
