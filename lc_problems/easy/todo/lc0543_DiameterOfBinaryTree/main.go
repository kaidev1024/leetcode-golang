type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := helper(root.Left)
		r := helper(root.Right)
		result = max(result, l+r)
		return 1 + max(l, r)
	}
	helper(root)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}