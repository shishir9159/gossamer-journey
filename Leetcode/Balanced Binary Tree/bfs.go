// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isBalanced(root *TreeNode) bool {
	order := []*TreeNode{}
	queue := []*TreeNode{root}
	heights := map[*TreeNode]int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}
		order = append(order, node)
		queue = append(queue, node.Left, node.Right)
	}

	for i := len(order) - 1; i >= 0; i-- {
		node := order[i]
		left, right := heights[node.Left], heights[node.Right]
		if abs(left-right) > 1 {
			return false
		}
		heights[node] = 1 + max(left, right)
	}

	return true
}
