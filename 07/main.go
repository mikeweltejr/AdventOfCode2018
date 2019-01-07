package main

func main() {
	firstSteps, lastSteps := newStepsFromFile("steps.txt")
	orderAllSteps(firstSteps, lastSteps)
	//fmt.Println(orderedSteps)
}
