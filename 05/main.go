package main

func main() {
	polymer := newPolymerFromFile("polymer.txt")
	checkForPolymer([]rune(polymer), 0)
}
