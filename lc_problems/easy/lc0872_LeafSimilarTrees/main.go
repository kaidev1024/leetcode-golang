
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	func getSequence(root *TreeNode) []int {
		if root == nil {
			return nil
		}

		result := make([]int, 0)

		var helper func(root *TreeNode)
		helper = func(root *TreeNode) {
			if root.Left == nil && root.Right == nil {
				result = append(result, root.Val)
				return
			}
			if root.Left != nil {
				helper(root.Left)
			}
			if root.Right != nil {
				helper(root.Right)
			}
		}
		helper(root)
		return result
	}

	func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
		a1 := getSequence(root1)
		a2 := getSequence(root2)
		if len(a1) != len(a2) {
			return false
		}
		for i, v := range a1 {
			if v != a2[i] {
				return false
			}
		}
		return true
	}