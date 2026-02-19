package submissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type removeElementInput struct {
	nums []int
	val  int
}
type removeElementTest struct {
	name     string
	input    removeElementInput
	expected int
}

var removeElementTests = []removeElementTest{
	{
		name: "ExampleTestCase1",
		input: removeElementInput{
			nums: []int{3, 2, 2, 3},
			val:  3,
		},
		expected: 2,
	},
	{
		name: "ExampleTestCase2",
		input: removeElementInput{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
		},
		expected: 5,
	},
}

func (test removeElementTest) run(t *testing.T) {

	res := removeElement(test.input.nums, test.input.val)

	assert.Equal(t, test.expected, res, "got %v want %v", res, test.expected)
}

func TestRemoveElement(t *testing.T) {

	for _, test := range removeElementTests {

		t.Run(test.name, test.run)

	}
}
