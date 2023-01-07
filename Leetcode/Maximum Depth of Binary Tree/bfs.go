// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type nodeDepth struct {
	node  *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maximumDepth := 0
	queue := []*nodeDepth{{node: root, depth: 1}}

	for len(queue) > 0 {
		node, depth := queue[0].node, queue[0].depth
		queue = queue[1:]

		maximumDepth = max(maximumDepth, depth)

		if node.Right != nil {
			queue = append(queue, &nodeDepth{node: node.Right, depth: 1 + depth})
		}

		if node.Left != nil {
			queue = append(queue, &nodeDepth{node: node.Left, depth: 1 + depth})
		}
	}

	return maximumDepth
}
