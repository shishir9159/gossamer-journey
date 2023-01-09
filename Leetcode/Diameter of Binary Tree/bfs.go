// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	depth := map[*TreeNode]int{}
	order := []*TreeNode{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		order = append(order, node)
		queue = append(queue, node.Left, node.Right)
	}

	for index := len(order) - 1; index >= 0; index-- {
		node := order[index]
		left, right := depth[node.Left], depth[node.Right]
		diameter = max(diameter, left+right)
		depth[node] = max(left, right) + 1
	}

	return diameter
}
