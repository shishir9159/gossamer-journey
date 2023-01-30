/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	for {
		if root.Left == nil {
			k--
			if k == 0 {
				return root.Val
			}
			root = root.Right
		} else {
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else {
				predecessor.Right = nil
				k--
				if k == 0 {
					return root.Val
				}
				root = root.Right
			}
		}
	}
}
