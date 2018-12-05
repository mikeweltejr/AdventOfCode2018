package main

import "fmt"

func main() {
	shifts := newShiftsFromFile("guards.txt")
	gShifts := shifts.buildGuardShiftList()
	guards := determineTotalMinutesSlept(gShifts)
	guardID := findGuardWithMostMinSlept(guards)
	minutes := countMinuteOfHighGuard(gShifts, guardID)
	highestMin := findMostFrequentMin(minutes)

	fmt.Println("GuardID:", guardID, "HighestMin:", highestMin)
	fmt.Println("Answer:", guardID*highestMin)
}
