// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func levelOrder(root *TreeNode) [][]int {
	order := [][]int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		} else if len(order) == depth {
			order = append(order, []int{})
		}

		order[depth] = append(order[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)
	return order
}
