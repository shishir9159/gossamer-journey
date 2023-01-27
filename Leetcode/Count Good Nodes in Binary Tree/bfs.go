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

	type nodeInfo struct {
		node   *TreeNode
		maxVal int
	}

	count, queue := 0, []nodeInfo{{root, root.Val}}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.maxVal <= top.node.Val {
			top.maxVal = top.node.Val
			count++
		}

		if top.node.Left != nil {
			queue = append(queue, nodeInfo{top.node.Left, top.maxVal})
		}

		if top.node.Right != nil {
			queue = append(queue, nodeInfo{top.node.Right, top.maxVal})
		}
	}

	return count
}
