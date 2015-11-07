package main

type step struct {
	value      int
	adjustment int
}

func getAdjustment(value int) int {
	switch value % 3 {
	case 1:
		return -1
	case 2:
		return 1
	default:
		return 0
	}
}

func getSteps(start int) []step {
	currentVal := start
	steps := []step{}
	for currentVal > 1 {
		adjustment := getAdjustment(currentVal)
		steps = append(steps, step{currentVal, adjustment})
		currentVal = (currentVal + adjustment) / 3
	}
	return append(steps, step{1, 0})
}

func main() {

}
