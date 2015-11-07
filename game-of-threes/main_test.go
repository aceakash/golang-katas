package main

import "testing"

type getStepsTestDataItem struct {
	start int
	steps []int
}

var getStepsTests = []getStepsTestDataItem{
	{1, []int{1}},
	{2, []int{2, 1}},
}

func TestGetSteps(t *testing.T) {
	for _, tt := range getStepsTests {
		actual := getSteps(tt.start)
		if len(actual) != len(tt.steps) {
			t.Errorf("getSteps %d has %d steps, expected %d", tt.start, len(actual), len(tt.steps))
		}
	}
}
