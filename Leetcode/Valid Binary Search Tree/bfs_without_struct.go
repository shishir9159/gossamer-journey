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
	}

	mins, maxs, queue := []int{math.MinInt64}, []int{math.MaxInt64}, []*TreeNode{root}

	for len(queue) > 0 {
		min, max, node := mins[0], maxs[0], queue[0]
		queue, mins, maxs = queue[1:], mins[1:], maxs[1:]

		if node.Val <= min || node.Val >= max {
			return false
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
			mins = append(mins, node.Val)
			maxs = append(maxs, max)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
			mins = append(mins, min)
			maxs = append(maxs, node.Val)
		}
	}

	return true
}
