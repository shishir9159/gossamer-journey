// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type nodeInfo struct {
	node    *TreeNode
	visited bool
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	depth := map[*TreeNode]int{}
	stack := []nodeInfo{{node: root}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			continue
		}

		if top.visited {
			left, right := depth[top.node.Left], depth[top.node.Right]
			diameter = max(diameter, left+right)
			depth[top.node] = max(left, right) + 1
		} else {
			stack = append(stack, nodeInfo{top.node, true},
				nodeInfo{node: top.node.Left}, nodeInfo{node: top.node.Right})
		}
	}

	return diameter
}
