
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	leftFound := lowestCommonAncestor(root.Left, p, q)
	rightFound := lowestCommonAncestor(root.Right, p, q)

	if leftFound != nil && rightFound != nil {
		return root
	}
	if leftFound != nil {
		return leftFound
	}
	if rightFound != nil {
		return rightFound
	}
	return nil
}