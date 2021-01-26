package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}

	if p.Val < root.Val && root.Val < q.Val {
		return root
	}

	if p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return lowestCommonAncestor(root.Left, p, q)

}
