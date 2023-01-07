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
	maximumDepth := 0
	stack := []*nodeDepth{{node: root, depth: 1}}

	for len(stack) > 0 {
		node, depth := stack[len(stack)-1].node, stack[len(stack)-1].depth
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		maximumDepth = max(maximumDepth, depth)

		stack = append(stack, &nodeDepth{node: node.Left, depth: 1 + depth}, &nodeDepth{node: node.Right, depth: 1 + depth})
	}

	return maximumDepth
}
