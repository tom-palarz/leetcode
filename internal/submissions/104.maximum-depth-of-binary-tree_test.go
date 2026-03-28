package submissions

import "testing"

func TestMaxDepthExample1(t *testing.T) {

	t.Parallel()

	// root = [3,9,20,nil,nil,15,7]
	root := new(TreeNode)
	root.Val = 3
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.Val = 9
	root.Right.Val = 20

	root.Left.Left = nil
	root.Left.Right = nil

	root.Right.Left = new(TreeNode)
	root.Right.Left.Val = 15
	root.Right.Left.Left = nil
	root.Right.Left.Right = nil

	root.Right.Right = new(TreeNode)
	root.Right.Right.Val = 7
	root.Right.Right.Left = nil
	root.Right.Right.Right = nil

	expect := 3

	ret := maxDepth(root)

	if ret != expect {
		t.Errorf("got %d, want %d", ret, expect)
	}
}

func TestMaxDepthExample2(t *testing.T) {

	t.Parallel()

	// root = [1,nil,2]
	root := new(TreeNode)
	root.Val = 1
	root.Left = nil
	root.Right = new(TreeNode)

	root.Right.Val = 2
	root.Right.Left = nil
	root.Right.Right = nil

	expect := 2

	ret := maxDepth(root)

	if ret != expect {
		t.Errorf("got %d, want %d", ret, expect)
	}
}
