package submissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mergeInput struct {
	nums1 []int
	nums2 []int
	m     int
	n     int
}
type mergeTest struct {
	name     string
	input    mergeInput
	expected []int
}

var mergeTests = []mergeTest{
	{
		name: "ExampleTestCase1",
		input: mergeInput{
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
		},
		expected: []int{1, 2, 2, 3, 5, 6},
	},
	{
		name: "ExampleTestCase2",
		input: mergeInput{
			nums1: []int{1},
			nums2: []int{},
			m:     1,
			n:     0,
		},
		expected: []int{1},
	},
	{
		name: "ExampleTestCase3",
		input: mergeInput{
			nums1: []int{0},
			nums2: []int{1},
			m:     0,
			n:     1,
		},
		expected: []int{1},
	},
	{
		name: "FirstSubmissionWrongAnswer33of59",
		input: mergeInput{
			nums1: []int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
			nums2: []int{1, 2, 2},
			m:     6,
			n:     3,
		},
		expected: []int{-1, 0, 0, 1, 2, 2, 3, 3, 3},
	},
	{
		name: "SecondSubmissionWrongAnswer50of59",
		input: mergeInput{
			nums1: []int{1, 0},
			nums2: []int{2},
			m:     1,
			n:     1,
		},
		expected: []int{1, 2},
	},
}

func (test mergeTest) run(t *testing.T) {

	res := make([]int, test.input.m+test.input.n)
	copy(res, test.input.nums1)

	merge(res, test.input.m, test.input.nums2, test.input.n)

	assert.Equal(t, test.expected, res, "got %v want %v", res, test.expected)
}

func TestMerge(t *testing.T) {

	for _, test := range mergeTests {

		t.Run(test.name, test.run)

	}
}
