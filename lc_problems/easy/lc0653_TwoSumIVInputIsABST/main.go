
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	func findTarget(root *TreeNode, k int) bool {
		m := make(map[int]struct{})

		var helper func(root *TreeNode) bool
		helper = func(root *TreeNode) bool {
			if root == nil {
				return false
			}
			if _, ok := m[k-root.Val]; ok {
				return true
			}
			m[root.Val] = struct{}{}
			return helper(root.Left) || helper(root.Right)
		}

		return helper(root)
	}