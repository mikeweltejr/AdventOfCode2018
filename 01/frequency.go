package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type frequency []int

func newFrequencyFromFile(filename string) frequency {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return convertStringArrayToIntArray(strings.Split(string(bs), "\r\n"))
}

func convertStringArrayToIntArray(stringArr []string) frequency {
	f := []int{}

	for _, s := range stringArr {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		f = append(f, int(i))
	}

	return f
}

func (f frequency) print() {
	for i, freq := range f {
		fmt.Println(i, freq)
	}
}

func (f frequency) sum() int {
	sum := 0
	for _, freq := range f {
		sum += freq
	}

	return sum
}

func (f frequency) checkForDuplicateFrequency(sum int, allFrequencies []int) {
	dup := 0
	dupFound := false

	fmt.Println("Calibrating...")

	for _, freq := range f {
		sum += freq
		for _, s := range allFrequencies {
			if s == sum {
				dup = s
				dupFound = true
				fmt.Println("Dup:", dup)
				break
			}
		}
		if dupFound {
			break
		}
		allFrequencies = append(allFrequencies, sum)
	}

	if !dupFound {
		f.checkForDuplicateFrequency(sum, allFrequencies)
	}
}
