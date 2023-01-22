func levelOrder(root *TreeNode) [][]int {
	order := [][]int{}
	stack, depths := []*TreeNode{root}, []int{0}
	for len(stack) > 0 {
		node, depth := stack[len(stack)-1], depths[len(stack)-1]
		stack, depths = stack[:len(stack)-1], depths[:len(stack)-1]

		if node == nil {
			continue
		} else if len(order) == depth {
			order = append(order, nil)
		}

		order[depth] = append(order[depth], node.Val)
		stack = append(stack, node.Right, node.Left)
		depths = append(depths, depth+1, depth+1)
	}

	return order
}
