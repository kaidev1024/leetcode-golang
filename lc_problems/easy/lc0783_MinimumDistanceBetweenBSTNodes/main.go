
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	minDiff := 10000000
	prev := -10000000

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}
		dif := root.Val - prev
		if dif < minDiff {
			minDiff = dif
		}
		prev = root.Val
		if root.Right != nil {
			helper(root.Right)
		}
	}
	helper(root)
	return minDiff
}