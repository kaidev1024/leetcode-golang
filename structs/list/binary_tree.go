package binaryTree

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func NewBinaryTreeNode(val int) *BinaryTreeNode {
	return &{Val: val}
}
