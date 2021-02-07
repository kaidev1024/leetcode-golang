
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	prev := -1
	diff := -1
	minDiff := -1

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}
		if prev == -1 {
			prev = root.Val
		} else if diff == -1 {
			diff = root.Val - prev
			minDiff = diff
			prev = root.Val
		} else {
			diff = root.Val - prev
			if diff < minDiff {
				minDiff = diff
			}
			prev = root.Val
		}
		if root.Right != nil {
			helper(root.Right)
		}
	}
	helper(root)
	return minDiff
}