package main

func main() {
	blankFreq := []int{}
	frequency := newFrequencyFromFile("frequency.txt")
	frequency.checkForDuplicateFrequency(0, blankFreq)
}
