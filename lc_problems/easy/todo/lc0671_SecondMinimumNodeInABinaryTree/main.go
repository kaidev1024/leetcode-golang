
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	func findSecondMinimumValue(root *TreeNode) int {
		return visit(root, root.Val)
	}

	func visit(root *TreeNode, target int) int {
		if root == nil {
			return -1
		}

		if root.Val > target {
			return root.Val
		}

		left := visit(root.Left, target)
		right := visit(root.Right, target)
		fmt.Println(left, right)
		if left == -1 {
			return right
		}

		if right == -1 {
			return left
		}

		return min(left, right)
	}

	func min(a, b int) int {
		if a < b {
			return a
		}

		return b
	}