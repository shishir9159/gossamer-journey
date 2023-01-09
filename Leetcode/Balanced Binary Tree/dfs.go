type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left, right := depth(node.Left), depth(node.Right)
	// if left < 0 || right < 0 || left-right > 1 || right-left > 1 {
	if left == -1 || right == -1 || math.Abs(float64(left-right)) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}
