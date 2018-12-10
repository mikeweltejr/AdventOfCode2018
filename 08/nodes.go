package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type nodes []int

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

	return nodeArr
}

func (n nodes) sum() int {
	sum := 0
	for _, num := range n {
		sum += num
	}

	return sum
}
