
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	result := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Val >= low && root.Val <= high {
			result += root.Val
			if root.Left != nil {
				helper(root.Left)
			}
			if root.Right != nil {
				helper(root.Right)
			}
		}
		if root.Val > high && root.Left != nil {
			helper(root.Left)
		}
		if root.Val < low && root.Right != nil {
			helper(root.Right)
		}
	}
	helper(root)
	return result
}