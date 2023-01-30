/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return nil
	}

	var order []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue, order = queue[1:], append(order, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	sort.Ints(order)
	return order[k-1]
}
