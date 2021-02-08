
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	result := &TreeNode{Val: root1.Val + root2.Val}
	result.Left = mergeTrees(root1.Left, root2.Left)
	result.Right = mergeTrees(root1.Right, root2.Right)
	return result
}