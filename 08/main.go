package main

import "fmt"

func main() {
	nodes := newNodesFromFile("nodes.txt")
	fmt.Println("Sum:", nodes.sum())
}
