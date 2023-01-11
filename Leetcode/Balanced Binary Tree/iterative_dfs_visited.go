func isBalanced(root *TreeNode) bool {
	type nodeInfo struct {
		node    *TreeNode
		visited bool
	}

	stack := []nodeInfo{{root, false}}
	heightStack := []int{}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			heightStack = append(heightStack, 0)
			continue
		} else if !top.visited {
			stack = append(stack, nodeInfo{top.node, true}, nodeInfo{top.node.Left, false}, nodeInfo{top.node.Right, false})
			continue
		}

		left, right := heightStack[len(heightStack)-2], heightStack[len(heightStack)-1]
		heightStack = heightStack[:len(heightStack)-2]
		if math.Abs(float64(left-right)) > 1 {
			return false
		}

		heightStack = append(heightStack, 1+max(left, right))
	}

	return true
}
