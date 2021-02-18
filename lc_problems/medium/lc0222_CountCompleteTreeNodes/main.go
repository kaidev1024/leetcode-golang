
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func power2(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * power2(n-1)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil {
		if root.Left == nil {
			return 1
		}
		return 2
	}

	l := root.Left
	lc := 1
	for l.Left != nil {
		lc++
		l = l.Left
	}

	r := root.Right

	lr := 1
	for r.Right != nil {
		lr++
		r = r.Right
	}
	if lc == lr {
		return power2(lc+1) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}