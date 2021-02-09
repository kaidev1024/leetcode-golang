
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, v int) (*TreeNode, int) {
	var parent *TreeNode
	depth := 0

	var traverse func(root, p *TreeNode, d, v int)
	traverse = func(root, p *TreeNode, d, v int) {
		if root == nil {
			return
		}

		if root.Val == v {
			parent = p
			depth = d
			return
		}
		traverse(root.Left, root, d+1, v)
		traverse(root.Right, root, d+1, v)
	}
	traverse(root, nil, 0, v)
	return parent, depth
}

func isCousins(root *TreeNode, x int, y int) bool {
	px, dx := helper(root, x)
	py, dy := helper(root, y)
	if px != py && dx == dy {
		return true
	}
	return false
}