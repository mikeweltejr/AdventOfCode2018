package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type elfStep struct {
	step rune
	prev *elfStep
	next *elfStep
}

func newStepsFromFile(filename string) ([]rune, []rune) {
	firstSteps := []rune{}
	lastSteps := []rune{}
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	allSteps := strings.Split(string(bs), "\r\n")

	for _, s := range allSteps {
		step1 := s[5:6]
		step2 := s[strings.LastIndex(s, "step")+5 : strings.LastIndex(s, "can")-1]

		firstSteps = append(firstSteps, rune(step1[0]))
		lastSteps = append(lastSteps, rune(step2[0]))
	}

	return firstSteps, lastSteps
}

func orderAllSteps(preSteps, postSteps []rune) {
	graph := NewGraph(len(preSteps) * 2)
	addedNodes := make(map[string]int)

	for i := 0; i < len(preSteps); i++ {
		strPre := strconv.QuoteRune(preSteps[i])
		strPost := strconv.QuoteRune(postSteps[i])
		if addedNodes[strPre] == 0 {
			addedNodes[strPre] = 1
			graph.AddNode(strPre)

		}
		if addedNodes[strPost] == 0 {
			addedNodes[strPost] = 1
			graph.AddNode(strPost)
		}
	}

	for i := 0; i < len(preSteps); i++ {
		graph.AddEdge(strconv.QuoteRune(preSteps[i]), strconv.QuoteRune(postSteps[i]))
	}

	result, ok := graph.Toposort()
	if !ok {
		panic("cycle detected")
	}

	print(result)
}

func print(s []string) {
	resultStr := ""
	for _, str := range s {
		resultStr += strings.Replace(str, "'", "", -1)
	}

	fmt.Println(resultStr)
}

func sort(pre, post []rune) ([]rune, []rune) {

	for i := 0; i < len(pre)-1; i++ {
		for j := i + 1; j < len(pre); j++ {
			if pre[i] > pre[j] {
				preTemp := pre[j]
				postTemp := post[j]
				pre[j] = pre[i]
				post[j] = post[i]
				pre[i] = preTemp
				post[i] = postTemp
			}
		}
	}

	return pre, post
}

// func orderSteps(firstSteps, lastSteps []rune) []elfStep {
// 	orderedList := []elfStep{}
// 	elfStep1 := elfStep{}
// 	elfStep2 := elfStep{}

// 	for i := 0; i < len(firstSteps); i++ {
// 		step1 := firstSteps[i]
// 		step2 := lastSteps[i]

// 		if len(orderedList) == 0 {
// 			elfStep2 = elfStep{step: step2}
// 			elfStep1 = elfStep{step: step1, next: &elfStep2}
// 			elfStep2.prev = &elfStep1
// 			orderedList = append(orderedList, elfStep1)
// 			orderedList = append(orderedList, elfStep2)
// 		} else {
// 			foundStep1 := findInList(orderedList, step1)
// 			currentNext := foundStep1.next
// 			fmt.Println(currentNext)
// 			elfStep2 = elfStep{step: step2, next: currentNext, prev: foundStep1}
// 			fmt.Println("elfStep2:", elfStep2)
// 			// foundStep2 := findInList(orderedList, step2)
// 			//fmt.Println("FoundStep1:", foundStep1, "FoundStep2:", foundStep2)

// 			// if step1 exists and step2 exist do nothing
// 			// if foundStep1.next == nil && foundStep1.prev == nil &&
// 			// 	foundStep2.next == nil && foundStep2.prev == nil {
// 			// 	continue
// 			// }
// 			// if (foundStep1.next != nil || foundStep1.prev != nil) &&
// 			// 	(foundStep2.next == nil && foundStep2.prev == nil) {

// 			// 	elfStep2 = elfStep{step: step2, next: foundStep1.next, prev: foundStep1}
// 			// 	//foundStep1.next.prev = &elfStep2
// 			// 	//foundStep1.next = &elfStep2

// 			orderedList = append(orderedList, elfStep2)
// 			// 	fmt.Println(orderedList)
// 			// }
// 			// if (foundStep1.next == nil && foundStep1.prev == nil) &&
// 			// 	(foundStep2.next != nil || foundStep2.prev != nil) {
// 			// 	currentNodePrev := foundStep2.prev
// 			// 	foundStep2.prev = &elfStep1

// 			// 	if currentNodePrev != nil {
// 			// 		currentNodePrev.next = &elfStep1
// 			// 		elfStep1.next = foundStep2
// 			// 	}

// 			// 	orderedList = append(orderedList, elfStep1)
// 			// }
// 			// if step1 does not exist and step2 does exist
// 			//		- step2 prev needs to be step1
// 			//		- step2 old prev next needs to be step1
// 			//		- step1 next needs to be step2
// 			// if neither exist
// 			//		- add step1 with next of step 2 prev of last item in list
// 			//		- add step2 with prev of step1 next nil
// 		}
// 	}

// 	printDetails(orderedList)
// 	return orderedList
// }

// func printDetails(e []elfStep) {
// 	for i := 0; i < len(e); i++ {
// 		item := e[i]
// 		fmt.Println("Step:", item.step, "Prev:", item.prev, "Next:", item.next)
// 	}
// }

// func printList(e []elfStep) {
// 	currentNode := findFirst(e)

// 	for currentNode != nil {
// 		fmt.Println("Step:", strconv.QuoteRune(currentNode.step))
// 		currentNode = currentNode.next
// 	}
// }

// func findFirst(e []elfStep) *elfStep {
// 	for i := 0; i < len(e); i++ {
// 		if e[i].prev == nil {
// 			return &e[i]
// 		}
// 	}

// 	return &e[0]
// }

// func findInList(e []elfStep, value rune) *elfStep {
// 	for i := 0; i < len(e); i++ {
// 		if e[i].step == value {
// 			return &e[i]
// 		}
// 	}

// 	return &elfStep{next: nil, prev: nil}
// }
