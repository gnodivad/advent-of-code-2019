package intcode

import (
	"gnodivad/aoc/intcode"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	inputs  []int
	outputs []int
}

func TestAdd(t *testing.T) {
	expected := []int{2, 0, 0, 0, 99}

	inputs := []int{1, 0, 0, 0, 99}

	intcode.Run(inputs)

	assert.ElementsMatch(t, expected, inputs, "Output and expected should be the same.")
}

func TestMultiply(t *testing.T) {
	expected := []int{2, 3, 0, 6, 99}

	inputs := []int{2, 3, 0, 3, 99}

	intcode.Run(inputs)

	assert.ElementsMatch(t, expected, inputs, "Output and expected should be the same.")
}
