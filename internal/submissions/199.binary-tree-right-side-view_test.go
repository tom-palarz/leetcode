package submissions

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRightSideViewExample1(t *testing.T) {

	// root = [1,2,3,nil,5,nil,4]
	root := new(TreeNode)

	root.Val = 1

	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Left.Left = nil

	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 5
	root.Left.Right.Left = nil
	root.Left.Right.Right = nil

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = nil
	root.Right.Right = new(TreeNode)

	root.Right.Right.Val = 4
	root.Right.Right.Left = nil
	root.Right.Right.Right = nil

	expected := []int{1, 3, 4}

	result := rightSideView(root)

	if !cmp.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestRightSideViewExample2(t *testing.T) {

	// root = [1,2,3,4,nil,nil,nil,5]
	root := new(TreeNode)

	root.Val = 1

	root.Left = new(TreeNode)
	root.Left.Val = 2
	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = nil
	root.Right.Right = nil

	root.Left.Right = nil
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 4
	root.Left.Left.Right = nil

	root.Left.Left.Left = new(TreeNode)
	root.Left.Left.Left.Val = 5
	root.Left.Left.Left.Right = nil
	root.Left.Left.Left.Left = nil

	expected := []int{1, 3, 4, 5}

	result := rightSideView(root)

	if !cmp.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestRightSideViewExample3(t *testing.T) {

	// root = 1,nil,3
	root := new(TreeNode)

	root.Val = 1

	root.Left = nil

	root.Right = new(TreeNode)
	root.Right.Val = 3
	root.Right.Left = nil
	root.Right.Right = nil

	expected := []int{1, 3}

	result := rightSideView(root)

	if !cmp.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}

func TestRightSideViewExample4(t *testing.T) {

	var root *TreeNode = nil

	expected := []int{}

	result := rightSideView(root)

	if !cmp.Equal(result, expected) {
		t.Errorf("got %v, want %v", result, expected)
	}
}
