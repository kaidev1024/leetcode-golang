
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	func findTilt(root *TreeNode) int {
		if root == nil {
			return 0
		}
		result := 0
		absDiff := func(a, b int) int {
			if a > b {
				return a - b
			}
			return b - a
		}

		var helper func(root *TreeNode) int
		helper = func(root *TreeNode) int {
			if root == nil {
				return 0
			}
			l := helper(root.Left)
			r := helper(root.Right)
			result += absDiff(l, r)
			return l + r + root.Val
		}
		helper(root)
		return result
	}