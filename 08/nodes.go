package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type nodes []int

var inputs = nodes{}
var index = 0
var total = 0

func newNodesFromFile(filename string) nodes {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	strNumbers := strings.Split(string(bs), " ")
	nodeArr := nodes{}

	for _, n := range strNumbers {
		intVal, err := strconv.Atoi(n)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		nodeArr = append(nodeArr, intVal)
	}

	inputs = nodeArr
	return nodeArr
}

func sumMetadata() int {
	quantities := inputs[index+1]
	metadata := inputs[index+1]
	sum := 0

	if quantities == 0 {
		for m := 0; m < metadata; m++ {
			i := inputs[index+1]
			sum += i
			total += i
		}
	} else {
		values := []int{}
		for s := 0; s < quantities; s++ {
			values = append(values, sumMetadata())
		}
		for m := 0; m < metadata; m++ {
			index++
			i := inputs[index+1]
			total += i
			if i >= 1 && i < len(values) {
				sum += values[i-1]
			}
		}
	}

	fmt.Println(sum)
	return sum

	// currentMetadataNode := len(n)
	// usedNodes := map[int]bool{}
	// sum := 0
	// fmt.Println("Len:", len(n))
	// for i := 0; i < len(n)-1; i++ {
	// 	if usedNodes[i] == true {
	// 		continue
	// 	}

	// 	usedNodes[i] = true
	// 	usedNodes[i+1] = true

	// 	quantity := n[i]
	// 	metadata := n[i+1]
	// 	subMetaDataNodes := 0

	// 	//fmt.Println("UsedNodes:", usedNodes)

	// 	if quantity > 0 {
	// 		fmt.Println("Quantity:", quantity, "Metadata:", metadata)
	// 		for j := metadata; j > 0; j-- {
	// 			fmt.Println("CurrentMetaNode:", currentMetadataNode, "j:", j, "Pos:", currentMetadataNode-j, "Val:", n[currentMetadataNode-j])
	// 			sum += n[currentMetadataNode-j]
	// 			usedNodes[currentMetadataNode-j] = true
	// 			subMetaDataNodes++
	// 		}
	// 	} else {
	// 		fmt.Println("Q 0", "Quantity:", quantity, "Metadata:", metadata)
	// 		for j := i + 2; j < i+2+metadata; j++ {
	// 			fmt.Println("N[j]:", n[j])
	// 			usedNodes[j] = true
	// 			sum += n[j]
	// 		}
	// 	}

	// 	fmt.Println("Sum:", sum)
	// 	currentMetadataNode = currentMetadataNode - subMetaDataNodes
	// }

	// //fmt.Println("Sum:", sum)
	// return sum
}

func printTotal() {
	fmt.Println("Total:", total)
}

func (n nodes) sum() int {
	sum := 0
	for _, num := range n {
		sum += num
	}

	return sum
}
