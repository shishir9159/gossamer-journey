// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	} else if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
