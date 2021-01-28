package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(tn *TreeNode) bool {
	if tn != nil && tn.Left == nil && tn.Right == nil {
		return true
	}
	return false
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var result int
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if isLeaf(root.Left) {
			result += root.Left.Val
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
