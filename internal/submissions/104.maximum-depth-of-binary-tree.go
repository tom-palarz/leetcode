package submissions

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	// DFS using recursion
	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)

	// return the greater of the right or left
	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}
