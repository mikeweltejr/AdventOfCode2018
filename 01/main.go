package main

import "fmt"

func main() {
	frequency := newFrequencyFromFile("frequency.txt")
	fmt.Println(frequency.sum())
}
