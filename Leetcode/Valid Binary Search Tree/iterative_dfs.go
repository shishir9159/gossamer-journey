func isValidBST(root *TreeNode) bool {
	type boundNode struct {
		min, max *int
		node     *TreeNode
	}

	stack := []boundNode{{root, nil, nil}}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := item.node
		if node == nil {
			continue
		} else if item.min != nil && node.Val <= *item.min {
			return false
		} else if item.max != nil && node.Val >= *item.max {
			return false
		}

		stack = append(stack, boundNode{&node.Val, node.Left, item.min}, boundNode{&node.Val, node.Right, item.max})
	}

	return true
}

func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		node.Left, node.Right = node.Right, node.Left
		stack = append(stack, node.Left, node.Right)
	}

	return root
}
