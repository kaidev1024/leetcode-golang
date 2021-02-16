
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			helper(root.Right)
		}
	}
	if root != nil {
		helper(root)
	}
	return result
}