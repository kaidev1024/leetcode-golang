
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	l := math.MinInt64
	r := math.MaxInt64

	var helper func(root *TreeNode, l, r int) bool
	helper = func(root *TreeNode, l, r int) bool {
		if root == nil {
			return true
		}
		if root.Val <= l || root.Val >= r {
			return false
		}
		return helper(root.Left, l, root.Val) && helper(root.Right, root.Val, r)
	}

	return helper(root, l, r)
}