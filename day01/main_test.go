package main

import (
	"testing"
)

var testCases = []struct {
	moduleWeight int
	expectedFuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestCalculateFuelRequired(t *testing.T) {
	for _, testCase := range testCases {
		result := CalculateFuelRequired(testCase.moduleWeight)
		if result != testCase.expectedFuel {
			t.Errorf("Error, expected %d fuel, got %d", testCase.expectedFuel, result)
		}
	}
}
