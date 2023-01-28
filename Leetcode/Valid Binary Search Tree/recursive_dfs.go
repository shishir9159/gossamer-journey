/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
	if node == nil {
		return true
	} else if min != nil && node.Val <= *min {
		return false
	} else if max != nil && node.Val >= *max {
		return false
	}

	return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
}
