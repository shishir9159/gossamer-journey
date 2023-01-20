// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil { // breaks template pattern
			continue
		}

		similar, pairs := true, [][2]*TreeNode{{node, subRoot}}
		for len(pairs) > 0 {
			a, b := pairs[0][0], pairs[0][1]
			pairs = pairs[1:]
			if a == nil && b == nil {
				continue
			} else if a == nil || b == nil || a.Val != b.Val {
				similar = false
				break
			}

			pairs = append(pairs, [2]*TreeNode{a.Left, b.Left}, [2]*TreeNode{a.Right, b.Right}) // breaks template pattern
		}

		if similar {
			return true
		}

		queue = append(queue, node.Left, node.Right)
	}

	return false
}
