/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	type nodeInfo struct {
		node  *TreeNode
		depth int
	}

	order, stack := [][]int{}, []nodeInfo{{root, 0}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.node == nil {
			continue
		} else if len(order) == top.depth {
			order = append(order, []int{})
		}

		order[top.depth] = append(order[top.depth], top.node.Val)
		stack = append(stack, nodeInfo{top.node.Right, top.depth + 1}, nodeInfo{top.node.Left, top.depth + 1})
	}

	return order
}
