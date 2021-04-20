type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var helper func(root *TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val >= low && root.Val <= high {
			root.Left = helper(root.Left)
			root.Right = helper(root.Right)
			return root
		}
		if root.Val < low {
			return helper(root.Right)
		}
		return helper(root.Left)
	}

	return helper(root)
}