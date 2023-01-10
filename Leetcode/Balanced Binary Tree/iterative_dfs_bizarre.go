// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isBalanced(root *TreeNode) bool {
	heights := make(map[*TreeNode]int)
	stack := make([]*TreeNode, 0, 16)
	var lastNode *TreeNode

	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		top := stack[len(stack)-1]
		// lastNode == top.Right is true when right node is already processed
		if top.Right != nil && lastNode != top.Right {
			node = top.Right
			continue
		}

		left, right := heights[top.Left], heights[top.Right]
		if math.Abs(float64(left-right)) > 1 {
			return false
		}
		heights[top] = 1 + max(left, right)
		lastNode, stack = top, stack[:len(stack)-1]
	}

	return true
}
