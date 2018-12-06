package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type shifts []string
type guardShift struct {
	id         int
	time       string
	minute     int
	action     int
	actionText string
}
type guardShifts []guardShift

func newShiftsFromFile(filename string) shifts {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), "\r\n")
}

func (s shifts) print() {
	for _, shift := range s {
		fmt.Println("Shift:", shift)
	}
}

func (s shifts) buildGuardShiftList() guardShifts {
	gShifts := guardShifts{}

	for _, shift := range s {
		charArr := []rune(shift)
		id := 0
		action := 0
		time := string(charArr[strings.Index(shift, "[")+1 : strings.Index(shift, "]")])
		minute, minErr := strconv.Atoi(time[strings.Index(time, ":")+1 : strings.Index(time, ":")+3])

		if minErr != nil {
			fmt.Println("Error:", minErr)
			os.Exit(1)
		}

		if strings.Index(shift, "begins") >= 0 {
			numID, err := strconv.Atoi(string(charArr[strings.Index(shift, "Guard")+7 : strings.Index(shift, "begins")-1]))

			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}

			id = numID
		}

		if strings.Contains(shift, "Guard") {
			action = 1
		}
		if strings.Contains(shift, "falls") {
			action = 2
		}
		if strings.Contains(shift, "wakes") {
			action = 3
		}

		gShifts = append(gShifts, guardShift{id: id, time: time, action: action, minute: minute})
	}

	gShifts.sort()
	g := addIds(gShifts)
	//g.print()
	return g
}

func addIds(g guardShifts) guardShifts {
	id := 0

	for i := 0; i < len(g); i++ {
		gShift := g[i]
		if gShift.id != 0 {
			id = gShift.id
		}
		if gShift.id == 0 {
			g[i].id = id
		}
	}

	return g
}

func (g guardShifts) sort() {
	sort.Slice(g, func(i, j int) bool {
		return g[i].time < g[j].time
	})
}

func (g guardShifts) print() {
	for i := 0; i < len(g); i++ {
		gShift := g[i]
		fmt.Println("ID:", gShift.id, "Time:", gShift.time, "Action:", gShift.action)
	}
}

func determineTotalMinutesSlept(g guardShifts) [5000]int {
	guards := [5000]int{}
	guardID := 0
	startMin := 0
	totalTimeSlept := 0

	for i := 0; i < len(g); i++ {
		gShift := g[i]
		if gShift.action == 1 {
			guardID = gShift.id
		}
		if gShift.action == 2 {
			startMin = gShift.minute
		}
		if gShift.action == 3 {
			totalTimeSlept = gShift.minute - startMin
			guards[guardID] += totalTimeSlept
		}
	}

	return guards
}

func findGuardWithMostMinSlept(guards [5000]int) int {
	highGuardID := 0
	highGuardCount := 0
	for i := 0; i < len(guards); i++ {
		if guards[i] > highGuardCount {
			highGuardCount = guards[i]
			highGuardID = i
		}
	}

	return highGuardID
}

func countMinuteOfHighGuard(g guardShifts, guardID int) [59]int {
	minutes := [59]int{}
	startMin := 0

	for i := 0; i < len(g); i++ {
		gShift := g[i]
		if gShift.id == guardID {
			guardID = gShift.id

			if gShift.action == 2 {
				startMin = gShift.minute
			}
			if gShift.action == 3 {
				for i := startMin; i < gShift.minute; i++ {
					minutes[i]++
				}
			}
		}
	}

	return minutes
}

func countHighMinForAllGuards(g guardShifts) {
	minutes := [5000][60]int{}
	startMin := 0
	guardID := 0

	for i := 0; i < len(g); i++ {
		gShift := g[i]
		if gShift.action == 1 {
			guardID = gShift.id
		}
		if gShift.action == 2 {
			startMin = gShift.minute
		}
		if gShift.action == 3 {
			for i := startMin; i < gShift.minute; i++ {
				minutes[guardID][i]++
			}
		}
	}

	highestMin := 0
	highestMinCount := 0

	for i := 0; i < len(minutes); i++ {
		for j := 0; j < len(minutes[i]); j++ {
			if minutes[i][j] > highestMinCount {
				highestMin = j
				highestMinCount = minutes[i][j]
				guardID = i
			}
		}
	}
	fmt.Println("GuardID:", guardID, "HighestMin:", highestMin)
	fmt.Println("Answer2:", guardID*highestMin)
	//return minutes
}

func findMostFrequentMin(minutes [59]int) int {
	highMin := 0
	highMinCount := 0

	for i := 0; i < len(minutes); i++ {
		if minutes[i] > highMinCount {
			highMinCount = minutes[i]
			highMin = i
		}
	}

	return highMin
}
