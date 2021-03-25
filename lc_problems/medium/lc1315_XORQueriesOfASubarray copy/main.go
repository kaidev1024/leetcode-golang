type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	result := 0
	var dfs func(root *TreeNode, gp, p bool)
	dfs = func(root *TreeNode, gp, p bool) {
		if gp {
			result += root.Val
		}
		if root.Left != nil {
			dfs(root.Left, p, root.Val%2 == 0)
		}
		if root.Right != nil {
			dfs(root.Right, p, root.Val%2 == 0)
		}
	}
	if root == nil {
		return 0
	}
	dfs(root, false, false)
	return result
}