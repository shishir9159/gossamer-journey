/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	order := []int{}

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var rightNode int
		// this doesn't work: for index := 0; index < len(queue); index++ {
		for _, node := range queue {
			queue = queue[1:]
			rightNode = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		order = append(order, rightNode)
	}

	return order
}
