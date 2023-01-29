func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type nodeInfo struct {
		node *TreeNode
		min  int
		max  int
	}

	queue := []*nodeInfo{{root, math.MinInt64, math.MaxInt64}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		min, max := item.min, item.max
		if item.node.Val <= min || item.node.Val >= max {
			return false
		}

		if item.node.Right != nil {
			queue = append(queue, &nodeInfo{item.node.Right, item.node.Val, max})
		}

		if item.node.Left != nil {
			queue = append(queue, &nodeInfo{item.node.Left, min, item.node.Val})
		}
	}

	return true
}
