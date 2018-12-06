package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
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
			fmt.Println("Answer:", "Length:", len(polymer))
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

func removeElements() {
	for i := 65; i < 91; i++ {
		polymer := newPolymerFromFile("polymer.txt")

		reg, err := regexp.Compile("[" + strconv.QuoteRune(rune(i)) + "" + strconv.QuoteRune(rune(i+32)) + "]")
		if err != nil {
			log.Fatal(err)
		}
		processedPolymer := reg.ReplaceAllString(polymer, "")

		fmt.Println("For Letters:", strconv.QuoteRune(rune(i)), strconv.QuoteRune(rune(i+32)))
		checkForPolymer([]rune(processedPolymer), 0)
	}
}

func removeUpperAndLower(polymer []rune, upperChar int, lowerChar int) []rune {
	fmt.Println(upperChar, lowerChar)
	for i := 0; i < len(polymer); i++ {
		if int(polymer[i]) == upperChar || int(polymer[i]) == lowerChar {
			fmt.Println("REMOVING")
			polymer = append(polymer[:i], polymer[i+1:]...)
		}
	}

	return polymer
}
