func serialize(root *TreeNode) string {
	if root == nil {
		return "$#"
	}

	return "$" + strconv.Itoa(root.Val) + serialize(root.Left) + serialize(root.Right)
}

func prefixFunction(s string) []int {
	n := len(s)
	pi := make([]int, n)

	for index := 1; index < n; index++ {
		length := pi[index-1]

		for length > 0 && s[index] != s[length] {
			length = pi[length-1]
		}

		if s[index] == s[length] {
			length++
		}

		pi[index] = length
	}

	return pi
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	serializedSubRoot := serialize(subRoot)
	combined := serializedSubRoot + "|" + serialize(root)

	piValues := prefixFunction(combined)
	subLen := len(serializedSubRoot)

	for i := subLen + 1; i < len(combined); i++ {
		if piValues[i] == subLen {
			return true
		}
	}

	return false
}
