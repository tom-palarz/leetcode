package submissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type removeDuplicate2Input struct {
	nums []int
}
type removeDuplicate2Test struct {
	name             string
	input            removeDuplicate2Input
	expected         int
	expectedDeepCopy []int
}

var removeDuplicate2Tests = []removeDuplicate2Test{
	{
		name: "ExampleTestCase1",
		input: removeDuplicate2Input{
			nums: []int{1, 1, 1, 2, 2, 3},
		},
		expectedDeepCopy: []int{1, 1, 2, 2, 3},
		expected:         5,
	},
	{
		name: "ExampleTestCase2",
		input: removeDuplicate2Input{
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		},
		expectedDeepCopy: []int{0, 0, 1, 1, 2, 3, 3},
		expected:         7,
	},
}

func (test removeDuplicate2Test) run(t *testing.T) {

	in := make([]int, len(test.input.nums))
	copy(in, test.input.nums)

	res := removeDuplicates2(in)

	assert.Equal(t, test.expected, res, "got %v want %v", res, test.expected)

	for i := range test.expectedDeepCopy {
		assert.Equal(t, test.expectedDeepCopy[i], in[i], "at index %d got %v want %v", i, in[i], test.expectedDeepCopy[i])
	}
}

func TestRemoveDuplicate2(t *testing.T) {

	for _, test := range removeDuplicate2Tests {

		t.Run(test.name, test.run)

	}
}
