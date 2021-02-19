
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	result := 0
	flag := false
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if flag {
			return
		}
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		count++
		if count == k {
			result = root.Val
			flag = true
			return
		}
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	return result
}