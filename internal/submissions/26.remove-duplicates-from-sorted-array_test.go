package submissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type removeDuplicateInput struct {
	nums []int
}
type removeDuplicateTest struct {
	name             string
	input            removeDuplicateInput
	expected         int
	expectedDeepCopy []int
}

var removeDuplicateTests = []removeDuplicateTest{
	{
		name: "ExampleTestCase1",
		input: removeDuplicateInput{
			nums: []int{1, 1, 2},
		},
		expectedDeepCopy: []int{1, 2},
		expected:         2,
	},
	{
		name: "ExampleTestCase2",
		input: removeDuplicateInput{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		},
		expectedDeepCopy: []int{0, 1, 2, 3, 4},
		expected:         5,
	},
}

func (test removeDuplicateTest) run(t *testing.T) {

	in := make([]int, len(test.input.nums))
	copy(in, test.input.nums)

	res := removeDuplicates(in)

	assert.Equal(t, test.expected, res, "got %v want %v", res, test.expected)

	for i := range test.expectedDeepCopy {
		assert.Equal(t, test.expectedDeepCopy[i], in[i], "at index %d got %v want %v", i, in[i], test.expectedDeepCopy[i])
	}
}

func TestRemoveDuplicate(t *testing.T) {

	for _, test := range removeDuplicateTests {

		t.Run(test.name, test.run)

	}
}
