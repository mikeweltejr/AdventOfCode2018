package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func newPolymerFromFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return string(bs)
}

func checkForPolymer(polymer []rune, pos int) {
	if pos < 0 {
		pos = 0
	}
	for i := pos; i < len(polymer); i++ {
		if i+1 >= len(polymer) {
			fmt.Println("Answer:", string(polymer), "Length:", len(polymer))
			break
		}

		if (unicode.ToLower(polymer[i]) == unicode.ToLower(polymer[i+1])) &&
			((unicode.IsLower(polymer[i]) && unicode.IsUpper(polymer[i+1])) ||
				(unicode.IsUpper(polymer[i]) && unicode.IsLower(polymer[i+1]))) {
			polymer = append(polymer[:i], polymer[i+1:]...)
			polymer = append(polymer[:i], polymer[i+1:]...)
			checkForPolymer(polymer, i-1)
			break
		}
	}
}
