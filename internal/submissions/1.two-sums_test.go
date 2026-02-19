package submissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type twoSumInput struct {
	nums   []int
	target int
}
type twoSumTest struct {
	name     string
	input    twoSumInput
	expected []int
}

var twoSumTests = []twoSumTest{
	{
		name: "ExampleTestCase1",
		input: twoSumInput{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		expected: []int{0, 1},
	},
	{
		name: "ExampleTestCase2",
		input: twoSumInput{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		expected: []int{1, 2},
	},
	{
		name: "ExampleTestCase3",
		input: twoSumInput{
			nums:   []int{3, 3},
			target: 6,
		},
		expected: []int{0, 1},
	},
	{
		name: "FirstSubmissionWrongAnswer60of63",
		input: twoSumInput{
			nums:   []int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1},
			target: 11,
		},
		expected: []int{5, 11},
	},
}

func (test twoSumTest) run(t *testing.T) {

	res := twoSum(test.input.nums, test.input.target)

	assert.Equal(t, test.expected, res, "got %v want %v", res, test.expected)
}

func TestTwoSum(t *testing.T) {

	for _, test := range twoSumTests {

		t.Run(test.name, test.run)

	}
}
