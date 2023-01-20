func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	} else if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
