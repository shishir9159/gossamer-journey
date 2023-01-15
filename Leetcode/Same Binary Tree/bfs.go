
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}

	for len(queue1) > 0 && len(queue2) > 0 {
		for i := 0; i < len(queue1); i++ {
			nodeP, nodeQ := queue1[0], queue2[0]
			queue1, queue2 = queue1[1:], queue2[1:]

			if nodeP == nil && nodeQ == nil {
				continue
			} else if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
				return false
			}

			queue1 = append(queue1, nodeP.Left, nodeP.Right)
			queue2 = append(queue2, nodeQ.Left, nodeQ.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}
