package main

import "fmt"

func main() {
	w := newWordsFromFile("words.txt")
	twoCount, threeCount := w.findRepeat()

	fmt.Println("Checksum:", twoCount*threeCount)

	firstWord, secondWord := w.findFabric()
	fmt.Println(firstWord)
	fmt.Println(secondWord)
	fmt.Println(removeDifferentChar(firstWord, secondWord))
}
