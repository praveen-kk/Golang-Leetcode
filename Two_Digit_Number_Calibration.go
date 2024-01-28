/*
Initial Thought :

create a Trie DS with the words to  1..9 map.

insert the map and iterate through each line and search for the first and last digits.

Time complexity: o(m * n^2)
Space Complexity: o(m*n)
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
func (t *Trie) digit_calibration(inputLine string) (digit int) {
	first := 0
	last := 0
	words := map[string]string{"1": "one", "2": "two", "3": "three", "4": "four", "5": "five", "6": "six", "7": "seven", "8": "eight", "9": "nine"}

	// Insert words into the trie
	for key, word := range words {
		t.Insert(word)
		t.Insert(key)
	}
	firstDigit := t.find_first_digit(inputLine)
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
	secondDigit := t.findSecondDigit(inputLine)
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

	digit = first*10 + last
	return digit

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
	//line := "bgnnvfsnbpx29vsjrlmgmsqthreeqxvclkhlv"
	//result := trie.digit_calibration(line)
	//fmt.Printf("digit: %d", result)

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
		line := scanner.Text()

		// Calculate the calibration value for the current line
		calibrationValue := trie.digit_calibration(line)
		fmt.Printf("numbers %d\n", calibrationValue)

		// Add the calibration value to the sum
		sum += calibrationValue
	}
	fmt.Printf("sum %d", sum)

}
