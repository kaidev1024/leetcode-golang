
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	var helper func(root *TreeNode, sum int)
	helper = func(root *TreeNode, sum int) {
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			result += sum
			return
		}
		if root.Left != nil {
			helper(root.Left, sum)
		}
		if root.Right != nil {
			helper(root.Right, sum)
		}
	}
	helper(root, 0)

	return result
}