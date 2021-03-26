type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	var helper func(root *TreeNode) *TreeNode
	helper = func(root *TreeNode) *TreeNode {
		if root.Left != nil {
			root.Left = helper(root.Left)
		}
		if root.Right != nil {
			root.Right = helper(root.Right)
		}
		if root.Val == target && root.Left == nil && root.Right == nil {
			return nil
		}
		return root
	}

	return helper(root)
}