func levelOrder(root *TreeNode) [][]int {
	depth, order := 0, [][]int{}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		} else if len(order) == depth {
			order = append(order, []int{})
		}

		order[depth] = append(order[depth], node.Val)
		stack = append(stack, node.Left, node.Right)
		depth++
	}

	return order
}
