/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	order := [][]int{}

	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}

	return root
}
