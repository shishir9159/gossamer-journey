/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	type boundNode struct {
		min, max *int
		node     *TreeNode
	}

	stack := []boundNode{{nil, nil, root}}
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := item.node
		if node == nil {
			continue
		} else if item.min != nil && node.Val <= *item.min {
			return false
		} else if item.max != nil && node.Val >= *item.max {
			return false
		}

		stack = append(stack, boundNode{item.min, &node.Val, node.Left}, boundNode{&node.Val, item.max, node.Right})
	}

	return true
}
