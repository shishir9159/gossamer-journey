/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var order []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || len(order) == k {
			return
		}

		dfs(node.Left)
		order = append(order, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return order[k-1]
}
