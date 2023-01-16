// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func serialize(root *TreeNode) string {
	if root == nil {
		return "$#"
	}

	return "$" + strconv.Itoa(root.Val) + serialize(root.Left) + serialize(root.Right)
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	left, right := 0, 0

	for index := 1; index < n; index++ {
		if index <= right {
			z[index] = min(right-index+1, z[index-left])
		}

		for index+z[index] < n && s[z[index]] == s[index+z[index]] {
			z[index]++
		}

		if index+z[index]-1 > right {
			left = index
			right = index + z[index] - 1
		}
	}

	return z
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	serializedSubRoot := serialize(subRoot)
	combined := serializedSubRoot + "|" + serialize(root)

	zValues := zFunction(combined)
	subLen := len(serializedSubRoot)

	for i := subLen + 1; i < len(combined); i++ {
		if zValues[i] == subLen {
			return true
		}
	}
	return false
}
