package submissions

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func rightSideView(root *TreeNode) []int {

	queue := []*TreeNode{}
	result := []int{}

	if root == nil {
		return result
	}

	// append the root first before processing
	queue = append(queue, root)

	// BFS
	for len(queue) > 0 {
		// The right side will be the right-most (last) in the queue
		// since we enqueue starting with the left child
		result = append(result, queue[len(queue)-1].Val)

		// length at this level of the tree
		l := len(queue)

		// for all nodes at this level
		for i := 0; i < l; i++ {

			node := queue[0]
			queue[0] = nil // signal the Go GC to collect

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
		}

	}

	return result
}
