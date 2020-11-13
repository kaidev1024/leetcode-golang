package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
