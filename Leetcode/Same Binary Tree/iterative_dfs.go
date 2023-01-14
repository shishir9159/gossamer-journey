type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	type Pair struct {
		first, second *TreeNode
	}

	stack := []Pair{{p, q}}

	for len(stack) > 0 {
		node1, node2 := stack[len(stack)-1].first, stack[len(stack)-1].second
		stack = stack[:len(stack)-1]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		stack = append(stack, Pair{node1.Right, node2.Right}, Pair{node1.Left, node2.Left})
	}

	return true
}
