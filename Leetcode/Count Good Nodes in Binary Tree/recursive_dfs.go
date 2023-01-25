func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxVal int) int
	dfs = func(node *TreeNode, maxVal int) int {
		var nodeCount int

		if node == nil {
			return 0
		} else if maxVal <= node.Val {
			nodeCount++
		}

		maxVal = max(maxVal, node.Val)

		return nodeCount + dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
	}

	return dfs(root, root.Val)
}
