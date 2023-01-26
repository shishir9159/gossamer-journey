func goodNodes(root *TreeNode) int {
	count, maxStack, stack := 0, []int{root.Val}, []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		maxVal := maxStack[len(maxStack)-1]
		maxStack = maxStack[:len(maxStack)-1]

		if node == nil {
			continue
		} else if maxVal <= node.Val {
			maxVal = node.Val
			count++
		}

		stack = append(stack, node.Left, node.Right)
		maxStack = append(maxStack, maxVal, maxVal)
	}

	return count
}
