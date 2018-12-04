package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type elfFabric []string

var graphPoints = [][]int{}
var graph = [1000][1000]string{}
var uniqueId int = 0

func newFabricFromFile(filename string) elfFabric {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\r\n")
}

func (e elfFabric) print() {
	for _, fabric := range e {
		fmt.Println("Fabric:", fabric)
	}
}

func calculateUniqueness(e elfFabric) int {
	uniqueID := 0

	for _, f := range e {
		fabric := strings.Split(f, " ")

		id, err0 := strconv.Atoi(strings.Split(fabric[0], "#")[1])
		if err0 != nil {
			fmt.Println("Error:", err0)
			os.Exit(1)
		}

		left, err1 := strconv.Atoi(strings.Split(fabric[2], ",")[0])

		if err1 != nil {
			fmt.Println("Error:", err1)
			os.Exit(1)
		}

		topStr := strings.Split(fabric[2], ",")[1]
		top, err2 := strconv.Atoi(topStr[:len(topStr)-1])
		if err2 != nil {
			fmt.Println("Error:", err2)
			os.Exit(1)
		}

		length, err3 := strconv.Atoi(strings.Split(fabric[3], "x")[0])
		if err3 != nil {
			fmt.Println("Error:", err3)
			os.Exit(1)
		}

		height, err4 := strconv.Atoi(strings.Split(fabric[3], "x")[1])
		if err4 != nil {
			fmt.Println("Error", err4)
			os.Exit(1)
		}

		uniqueID = getUniqueID(left, top, length, height, id)

		if uniqueID != 0 {
			fmt.Println("UniqueId:", uniqueID)
			return uniqueID
		}
	}

	return 0
}

func (e elfFabric) getSharedSquareInches() int {
	graph = createGraph()

	for _, f := range e {
		fabric := strings.Split(f, " ")

		id, err0 := strconv.Atoi(strings.Split(fabric[0], "#")[1])
		if err0 != nil {
			fmt.Println("Error:", err0)
			os.Exit(1)
		}

		left, err1 := strconv.Atoi(strings.Split(fabric[2], ",")[0])

		if err1 != nil {
			fmt.Println("Error:", err1)
			os.Exit(1)
		}

		topStr := strings.Split(fabric[2], ",")[1]
		top, err2 := strconv.Atoi(topStr[:len(topStr)-1])
		if err2 != nil {
			fmt.Println("Error:", err2)
			os.Exit(1)
		}

		length, err3 := strconv.Atoi(strings.Split(fabric[3], "x")[0])
		if err3 != nil {
			fmt.Println("Error:", err3)
			os.Exit(1)
		}

		height, err4 := strconv.Atoi(strings.Split(fabric[3], "x")[1])
		if err4 != nil {
			fmt.Println("Error", err4)
			os.Exit(1)
		}

		calculatePoints(left, top, length, height, id)
	}

	commonPoints := totalSquareInchesInCommon(graphPoints)
	calculateUniqueness(e)
	return commonPoints
}

func calculatePoints(left int, top int, length int, height int, id int) [][]int {
	for i := 0; i < length; i++ {
		x := []int{}
		y := []int{}
		x = append(x, i+left)
		x = append(x, id)
		for j := 0; j < height; j++ {
			y = append(y, j+top)
		}
		graphPoints = append(graphPoints, x)
		graphPoints = append(graphPoints, y)
	}

	return graphPoints
}

func totalSquareInchesInCommon(points [][]int) int {
	count := 0
	for i := 0; i < len(points); i++ {
		if i%2 == 0 {
			x := points[i][0]
			yArr := points[i+1]
			for j := 0; j < len(yArr); j++ {
				y := yArr[j]
				if graph[x][y] == "#" {
					graph[x][y] = "X"
					count++
				}
				if graph[x][y] == "*" {
					graph[x][y] = "#"
				}
			}
		}
	}
	return count
}

func getUniqueID(left int, top int, length int, height int, id int) int {
	totalLength := length * height
	uniqueCount := 0
	for i := 0; i < length; i++ {
		for j := 0; j < height; j++ {
			if graph[left+i][j+top] == "#" {
				uniqueCount++
			}
		}
		if uniqueCount == totalLength {
			return id
		}
	}

	return 0
}

func createGraph() [1000][1000]string {
	g := [1000][1000]string{}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			g[i][j] = "*"
		}
	}

	return g
}

// func totalSquareInchesInCommon(points []graph) int {
// 	for i := 0; i < len(points); i++ {
// 		fmt.Println("i:", i)
// 		for j := i + 1; j < len(points); j++ {
// 			if reflect.DeepEqual(points[i], points[j]) {
// 				if !pointExists(overlappedPoints, points[i]) {
// 					overlappedPoints = append(overlappedPoints, points[i])
// 					fmt.Println(len(overlappedPoints))
// 				}
// 			}
// 		}
// 	}

// 	return len(overlappedPoints)
// }

// func pointExists(points []graph, point graph) bool {
// 	for i := 0; i < len(points); i++ {
// 		p := points[i]
// 		if reflect.DeepEqual(p, point) {
// 			return true
// 		}
// 	}

// 	return false
// }
