// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
//

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	depth := map[*TreeNode]int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node == nil {
			stack = stack[:len(stack)-1]
			continue
		}

		left, leftOk := depth[node.Left]
		right, rightOk := depth[node.Right]

		if (node.Left == nil || leftOk) && (node.Right == nil || rightOk) {
			stack = stack[:len(stack)-1]
			depth[node] = 1 + max(left, right)
			diameter = max(diameter, left+right)
		} else {
			stack = append(stack, node.Left, node.Right)
		}
	}

	return diameter
}
