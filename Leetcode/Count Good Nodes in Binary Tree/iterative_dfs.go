/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	type nodeInfo struct {
		node   *TreeNode
		maxVal int
	}

	count, stack := 0, []nodeInfo{{root, root.Val}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			continue
		} else if top.maxVal <= top.node.Val {
			top.maxVal = top.node.Val
			count++
		}

		stack = append(stack,
			nodeInfo{top.node.Left, top.maxVal},
			nodeInfo{top.node.Right, top.maxVal},
		)
	}

	return count
}
