package main

import (
	"fmt"
	"reflect"
	"testing"
)

type getStepsTestDataItem struct {
	start int
	steps []step
}

func TestGetSteps(t *testing.T) {
	var getStepsTests = []getStepsTestDataItem{
		{1, []step{
			{1, 0},
		}},
		{2, []step{
			{2, 1},
			{1, 0},
		}},
		{20, []step{
			{20, 1},
			{7, -1},
			{2, 1},
			{1, 0},
		}},
	}
	fmt.Println(getStepsTests)
	for _, tt := range getStepsTests {
		actual := getSteps(tt.start)
		if !reflect.DeepEqual(actual, tt.steps) {
			t.Errorf("getSteps %d was %v, expected %v", tt.start, actual, tt.steps)
		}
	}
}
