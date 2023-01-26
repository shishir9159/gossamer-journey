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

	count, maxQueue, queue := 0, []int{root.Val}, []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		maxVal := maxQueue[0]
		maxQueue = maxQueue[1:]

		if maxVal <= node.Val {
			maxVal = node.Val
			count++
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			maxQueue = append(maxQueue, maxVal)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			maxQueue = append(maxQueue, maxVal)
		}
	}

	return count
}
