// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type nodeInfo struct {
	node        *TreeNode
	depth       int
	diameter    int
	left, right *nodeInfo
}

func (node *nodeInfo) measurements() (depth, diameter int) {
	if node == nil {
		return 0, 0
	}

	return node.depth, node.diameter
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rootInfo := &nodeInfo{node: root}
	stack := []*nodeInfo{rootInfo}

	for len(stack) > 0 {
		top := stack[len(stack)-1]

		// left and right could be appended on each iteration, but this is more compact
		// if (top.node.Left != nil && top.left == nil) || (top.node.Right != nil && top.right == nil) {
		if top.node.Left != nil && top.left == nil {
			top.left = &nodeInfo{node: top.node.Left}
			stack = append(stack, top.left)
		} else if top.node.Right != nil && top.right == nil {
			top.right = &nodeInfo{node: top.node.Right}
			stack = append(stack, top.right)
		} else {
			stack = stack[:len(stack)-1]

			leftHeight, leftDiameter := top.left.measurements()
			rightHeight, rightDiameter := top.right.measurements()

			top.depth = 1 + max(leftHeight, rightHeight)
			top.diameter = max(leftHeight+rightHeight, max(leftDiameter, rightDiameter))
		}
	}

	return rootInfo.diameter
}
