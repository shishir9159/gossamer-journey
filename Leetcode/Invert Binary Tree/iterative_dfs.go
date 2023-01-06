// type TreeNode struct {
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
// }

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
