package main

import "fmt"

func (w words) findFabric() (string, string) {
	count := 0
	firstWord := ""
	secondWord := ""

	for i := 0; i < len(w); i++ {
		count, firstWord, secondWord = countDifferences(w, i+1, w[i])
		fmt.Println("Count:", count)
		if count == 1 {
			return firstWord, secondWord
		}
	}

	return firstWord, secondWord
}

func countDifferences(words []string, pos int, word string) (int, string, string) {
	count := 0
	diffWord := ""
	for i := pos; i < len(words); i++ {
		if count == 1 {
			break
		}
		count = 0
		diffWord = words[i]
		for j := 0; j < len(diffWord); j++ {
			if len(diffWord) != len(word) {
				break
			}

			if diffWord[j] != word[j] {
				if count > 1 {
					break
				}
				count++
			}
		}
	}

	return count, word, diffWord
}

func removeDifferentChar(word1 string, word2 string) string {
	retWord := ""
	for i := 0; i < len(word1); i++ {
		if word1[i] == word2[i] {
			retWord += string(word1[i])
		}
	}

	return retWord
}
