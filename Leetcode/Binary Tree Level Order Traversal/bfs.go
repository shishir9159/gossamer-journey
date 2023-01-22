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
		levelNodes := []int{}

		for index := 0; index < len(queue); index++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}

		order = append(order, levelNodes)
	}

	return order
}
