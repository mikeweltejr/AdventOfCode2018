package main

import "fmt"

func main() {
	fabrics := newFabricFromFile("fabric.txt")
	fmt.Println(fabrics.getSharedSquareInches())
}
