// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		left, right := depth(node.Left), depth(node.Right)
		maxDiameter = max(maxDiameter, left+right)
		return 1 + max(left, right)
	}

	depth(root)
	return maxDiameter
}
