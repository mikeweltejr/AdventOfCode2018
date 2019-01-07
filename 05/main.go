package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()

	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))

	polymer := newPolymerFromFile("polymer.txt")
	checkForPolymer([]rune(polymer), 0)
	//removeElements()

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
