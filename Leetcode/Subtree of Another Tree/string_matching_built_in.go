// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var serialize func(*TreeNode) string
	serialize = func(n *TreeNode) string {
		if n == nil {
			return "#"
		}

		return fmt.Sprintf("^%d%s%s", n.Val, serialize(n.Left), serialize(n.Right))
	}

	return strings.Contains(serialize(root), serialize(subRoot))
}
