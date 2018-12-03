package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type words []string

func newWordsFromFile(filename string) words {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\r\n")
}

func (w words) findRepeat() (int, int) {
	twoRepeatCount := 0
	threeRepeatCount := 0

	for _, word := range w {

		i, j := findLetterRepeats(word)

		if i {
			twoRepeatCount++
		}
		if j {
			threeRepeatCount++
		}
	}

	return twoRepeatCount, threeRepeatCount
}

func findLetterRepeats(word string) (bool, bool) {
	count := 0
	twoRepeats := false
	threeRepeats := false
	lettersChecked := []rune{}

	for _, letter := range word {
		if !hasLetterBeenChecked(lettersChecked, letter) {
			lettersChecked = append(lettersChecked, letter)
			count = getCountInString(word, letter)
			if count == 2 {
				twoRepeats = true
			}
			if count == 3 {
				threeRepeats = true
			}
		}
	}

	return twoRepeats, threeRepeats
}

func getCountInString(word string, letter rune) int {
	count := 0
	for _, l := range word {
		if l == letter {
			count++
		}
	}

	return count
}

func hasLetterBeenChecked(letters []rune, letter rune) bool {
	for _, l := range letters {
		if l == letter {
			return true
		}
	}

	return false
}
