func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if q.Val < p.Val {
		p, q = q, p
	}

	if root.Val <= q.Val && p.Val <= root.Val {
		return root
		// } else if root.Val < q.Val && root.Val < p.Val {
	} else if root.Val <= q.Val && root.Val <= p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return lowestCommonAncestor(root.Left, p, q)
}
