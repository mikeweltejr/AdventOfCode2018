package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func newCoordinatesFromFile(filename string) ([]int, []int) {
	bs, err := ioutil.ReadFile(filename)
	x := []int{}
	y := []int{}

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	coordinateStrings := strings.Split(string(bs), "\r\n")

	for _, cord := range coordinateStrings {
		c := strings.Split(cord, ", ")
		xInt, xErr := strconv.Atoi(c[0])
		yInt, yErr := strconv.Atoi(c[1])

		if xErr != nil {
			fmt.Println("Error:", xErr)
			os.Exit(1)
		}
		if yErr != nil {
			fmt.Println("Error:", yErr)
			os.Exit(1)
		}

		x = append(x, xInt)
		y = append(y, yInt)
	}

	return x, y
}

func calculateDistance(xCord, yCord []int) {
	maxX := getMaxInt(xCord)
	maxY := getMaxInt(yCord)
	//safeCount := 0

	grid := make([][]int, maxX+1)
	regions := make(map[int]int)

	for i := range grid {
		grid[i] = make([]int, maxY+1)
	}

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			best := maxX + maxY
			bestNumber := -1

			for i := 0; i < len(xCord); i++ {
				pointX := xCord[i]
				pointY := yCord[i]

				dist := int(math.Abs(float64(x-pointX)) + math.Abs(float64(y-pointY)))

				if dist < best {
					best = dist
					bestNumber = i
				} else if dist == best {
					bestNumber = -1
				}
			}

			grid[x][y] = bestNumber
			total := regions[bestNumber]
			total++
			regions[bestNumber] = total
		}
	}

	for x := 0; x <= maxX; x++ {
		bad := grid[x][0]
		delete(regions, bad)
		bad = grid[x][maxY]
		delete(regions, bad)
	}
	for y := 0; y <= maxY; y++ {
		bad := grid[0][y]
		delete(regions, bad)
		bad = grid[maxX][y]
		delete(regions, bad)
	}

	biggest := 0
	for _, v := range regions {
		if v > biggest {
			biggest = v
		}
	}

	fmt.Println("Regions:", regions)
	fmt.Println("Biggest:", biggest)

	inArea := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			size := 0
			for i := 0; i < len(xCord); i++ {
				pointX := xCord[i]
				pointY := yCord[i]
				dist := int(math.Abs(float64(x-pointX)) + math.Abs(float64(y-pointY)))
				size += dist
			}

			if size < 10000 {
				inArea++
			}
		}
	}

	fmt.Println("Area Size:", inArea)

}

func getMaxInt(intArr []int) int {
	maxValue := intArr[0]
	for _, i := range intArr {
		if i > maxValue {
			maxValue = i
		}
	}
	return maxValue
}
