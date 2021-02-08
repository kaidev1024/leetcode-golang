
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	result := strconv.Itoa(t.Val)

	if t.Left == nil && t.Right == nil {
		return result
	}
	if t.Left == nil && t.Right != nil {
		l := "()"
		r := "(" + tree2str(t.Right) + ")"
		return result + l + r
	}
	l := "(" + tree2str(t.Left) + ")"
	r := ""
	if t.Right != nil {
		r = "(" + tree2str(t.Right) + ")"
	}
	return result + l + r
}