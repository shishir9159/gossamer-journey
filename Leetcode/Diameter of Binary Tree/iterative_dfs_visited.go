// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func diameterOfBinaryTree(root *TreeNode) int {
	type nodeInfo struct {
		node    *TreeNode
		visited bool
	}

	diameter := 0
	heightStack := []int{}
	stack := []nodeInfo{{root, false}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			heightStack = append(heightStack, 0)
			continue
		} else if !top.visited {
			stack = append(stack, nodeInfo{top.node, true}, nodeInfo{top.node.Left, false}, nodeInfo{top.node.Right, false})
			continue
		}

		left, right := heightStack[len(heightStack)-2], heightStack[len(heightStack)-1]
		heightStack = heightStack[:len(heightStack)-2]
		diameter = max(diameter, left+right) // number of edges, not node

		heightStack = append(heightStack, 1+max(left, right))
	}

	return diameter
}
