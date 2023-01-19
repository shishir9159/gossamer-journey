// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}}
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		queue = queue[1:]

		if a == nil && b == nil {
			continue
		} else if a == nil || b == nil || a.Val != b.Val {
			return false
		}

		queue = append(queue, [2]*TreeNode{a.Left, b.Left}, [2]*TreeNode{a.Right, b.Right})
	}

	return true
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		} else if isSameTree(node, subRoot) {
			return true
		}

		queue = append(queue, node.Left, node.Right)
	}

	return false
}
