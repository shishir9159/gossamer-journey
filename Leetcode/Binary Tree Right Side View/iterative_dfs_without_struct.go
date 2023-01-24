/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	order, stack, depths := []int{}, []*TreeNode{root}, []int{0}
	for len(stack) > 0 {
		node, depth := stack[len(stack)-1], depths[len(stack)-1]
		stack, depths = stack[:len(stack)-1], depths[:len(stack)-1]

		if node == nil {
			continue
		} else if len(order) == depth {
			order = append(order, node.Val)
		}

		stack, depths = append(stack, node.Left, node.Right), append(depths, depth+1, depth+1)
	}

	return order
}
