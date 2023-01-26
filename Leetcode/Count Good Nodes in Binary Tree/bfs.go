/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type Item struct {
		node   *TreeNode
		maxVal int
	}

	queue := []Item{{root, root.Val}}
	count := 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		node := item.node
		maxVal := item.maxVal

		if maxVal <= node.Val {
			maxVal = node.Val
			count++
		}

		if node.Left != nil {
			queue = append(queue, Item{node.Left, maxVal})
		}

		if node.Right != nil {
			queue = append(queue, Item{node.Right, maxVal})
		}
	}

	return count
}
