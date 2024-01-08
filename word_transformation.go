/*
Given a dictionary of words, determine whether it is possible to transform a given word
into another with a fixed number of characters
isTransformable(start: string, end: string, dictionary: string[]): Boolean
dog
hat
dictionary = {'dot', 'cat', 'hot', 'hog', 'eat', 'dug', 'dig'}

dog => dot => hot => hat

Initial Thoughts :-

start node - dog
step 1 check for word with 1 letter difference to its next node-word in the dictionary
.. check every adj node matches this logic with the dest
check for the shortest path to reach the dest node - hat


*/

package main

import "fmt"

func isTransformable(start, end string, dictionary []string) bool {

	visited := make(map[string]bool)
	queue := []string{start}

	if len(start) != len(end) {
		return false
	}

	for len(queue) > 0 {
		current_word := queue[0]
		queue = queue[1:]

		if check_char_diff(current_word, end) {
			return true
		}
		for _, word := range dictionary {
			if !visited[word] && check_char_diff(current_word, word) {
				visited[word] = true
				queue = append(queue, word)
			}
		}

	}
	return false
}

func check_char_diff(word1, word2 string) bool {
	diffcount := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diffcount++
			if diffcount > 1 {
				return false
			}
		}
	}

	return diffcount == 1

}

func main() {

	dictionary := []string{"dot", "cat", "hot", "hog", "eat", "dug", "dig"}

	start := "dog"
	end := "hat"

	result := isTransformable(start, end, dictionary)
	fmt.Println(result)

}
