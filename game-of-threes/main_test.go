package main

import (
	"reflect"
	"testing"
)

var getStepsTests = []struct {
	start int
	steps []step
}{
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
	{100, []step{
		{100, -1},
		{33, 0},
		{11, 1},
		{4, -1},
		{1, 0},
	}},
}

func TestGetSteps(t *testing.T) {
	for _, tt := range getStepsTests {
		actual := getSteps(tt.start)
		if !reflect.DeepEqual(actual, tt.steps) {
			t.Errorf("getSteps %d was %v, expected %v", tt.start, actual, tt.steps)
		}
	}
}

var getAdjustmentTests = []struct{ v, a int }{
	{1, -1},
	{2, 1},
	{3, 0},
	{4, -1},
	{5, 1},
	{6, 0},
	{7, -1},
	{8, 1},
	{9, 0},
}

func TestGetAdjustment(t *testing.T) {
	for _, tt := range getAdjustmentTests {
		actual := getAdjustment(tt.v)
		if actual != tt.a {
			t.Errorf("getAdjustment %d was %d, expected %d", tt.v, actual, tt.a)
		}
	}
}
