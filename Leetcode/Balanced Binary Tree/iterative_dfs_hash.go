// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isBalanced(root *TreeNode) bool {
	stack := []*TreeNode{root}
	heights := map[*TreeNode]int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		left, right := heights[node.Left], heights[node.Right]
		if (node.Left != nil && left == 0) || (node.Right != nil && right == 0) {
			stack = append(stack, node, node.Left, node.Right) // different from template
			continue
		}

		if math.Abs(float64(left-right)) > 1 {
			return false
		}

		heights[node] = 1 + max(left, right)
	}

	return true
}
