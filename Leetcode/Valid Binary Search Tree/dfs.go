/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} else if root.Left != nil && root.Left.Val >= root.Val {
		return false
	} else if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return isValidBST(root.Left) && isValidBST(root.Right)
}
