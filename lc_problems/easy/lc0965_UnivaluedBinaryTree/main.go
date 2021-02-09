
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	v := root.Val
	r := true
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Val != v {
			r = false
			return
		}
		if root.Left != nil {
			helper(root.Left)
		}
		if root.Right != nil {
			helper(root.Right)
		}
	}
	if root.Left != nil {
		helper(root.Left)
	}
	if root.Right != nil {
		helper(root.Right)
	}
	return r
}