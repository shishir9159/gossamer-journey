/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	type Item struct {
		node   *TreeNode
		maxVal int
	}

	count, stack := 0, []Item{{root, root.Val}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node, maxVal := item.node, item.maxVal

		if node == nil {
			continue
		} else if maxVal <= node.Val {
			maxVal = node.Val
			count++
		}

		stack = append(stack,
			Item{node.Left, maxVal},
			Item{node.Right, maxVal},
		)
	}

	return count
}
