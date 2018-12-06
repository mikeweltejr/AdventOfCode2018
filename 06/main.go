package main

import "fmt"

func main() {
	x, y := newCoordinatesFromFile("coordinates.txt")
	fmt.Println("X", x, "Y", y)
	calculateDistance(x, y)
}
