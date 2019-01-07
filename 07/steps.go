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
