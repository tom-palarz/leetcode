package submissions

import "testing"

func TestIsSameTreeExample1(t *testing.T) {

	// p = [1,2,3], q = [1,2,3]
	pRoot := new(TreeNode)
	qRoot := new(TreeNode)

	pRoot.Val = 1
	qRoot.Val = 1

	pRoot.Left = new(TreeNode)
	pRoot.Left.Val = 2
	pRoot.Left.Left = nil
	pRoot.Left.Right = nil

	qRoot.Left = new(TreeNode)
	qRoot.Left.Val = 2
	qRoot.Left.Left = nil
	qRoot.Left.Right = nil

	pRoot.Right = new(TreeNode)
	pRoot.Right.Val = 3
	pRoot.Right.Left = nil
	pRoot.Right.Right = nil

	qRoot.Right = new(TreeNode)
	qRoot.Right.Val = 3
	qRoot.Right.Left = nil
	qRoot.Right.Right = nil

	expected := true
	actual := isSameTree(pRoot, qRoot)

	if actual != expected {
		t.Errorf("isSameTree(%v, %v) = %v, want %v", pRoot, qRoot, actual, expected)
	}

}

func TestIsSameTreeExample2(t *testing.T) {

	// p = [1,2], q = [1,nil,2]
	pRoot := new(TreeNode)
	qRoot := new(TreeNode)

	pRoot.Val = 1
	pRoot.Left = new(TreeNode)
	pRoot.Left.Val = 2
	pRoot.Left.Left = nil
	pRoot.Left.Right = nil
	pRoot.Right = nil

	qRoot.Val = 1
	qRoot.Right = new(TreeNode)
	qRoot.Right.Val = 2
	qRoot.Right.Left = nil
	qRoot.Right.Right = nil
	pRoot.Left = nil

	expected := false
	actual := isSameTree(pRoot, qRoot)

	if actual != expected {
		t.Errorf("isSameTree(%v, %v) = %v, want %v", pRoot, qRoot, actual, expected)
	}
}

func TestIsSameTreeExample3(t *testing.T) {

	// p = [1,2,1], q = [1,1,2]
	pRoot := new(TreeNode)
	qRoot := new(TreeNode)

	pRoot.Val = 1

	pRoot.Left = new(TreeNode)
	pRoot.Left.Val = 2
	pRoot.Left.Left = nil
	pRoot.Left.Right = nil

	pRoot.Right = new(TreeNode)
	pRoot.Right.Val = 1
	pRoot.Right.Left = nil
	pRoot.Right.Right = nil

	qRoot.Val = 1

	qRoot.Left = new(TreeNode)
	qRoot.Left.Val = 1
	qRoot.Left.Left = nil
	qRoot.Left.Right = nil

	qRoot.Right = new(TreeNode)
	qRoot.Right.Val = 2
	qRoot.Right.Left = nil
	qRoot.Right.Right = nil

	expected := false
	actual := isSameTree(pRoot, qRoot)

	if actual != expected {
		t.Errorf("isSameTree(%v, %v) = %v, want %v", pRoot, qRoot, actual, expected)
	}
}
