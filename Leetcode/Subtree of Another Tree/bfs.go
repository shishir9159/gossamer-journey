// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := [][2]*TreeNode{{p, q}}
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]

		a, b := pair[0], pair[1]
		if a == nil && b == nil {
			continue
		} else if a == nil || b == nil || a.Val != b.Val {
			return false
		}

		queue = append(queue, [2]*TreeNode{a.Left, a.Right}, [2]*TreeNode{b.Left, b.Right})
	}
	return true
}
