/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		newRoot := &TreeNode{v, nil, nil}
		newRoot.Left = root
		return newRoot
	}
	try(root, 1, d, v)
	return root
}

func try(root *TreeNode, i int, d int, v int) {
	if root == nil {
		return
	}
	if i == d-1 {
		tmpL := root.Left
		root.Left = &TreeNode{v, nil, nil}
		root.Left.Left = tmpL

		tmpR := root.Right
		root.Right = &TreeNode{v, nil, nil}
		root.Right.Right = tmpR
		return
	}
	try(root.Left, i+1, d, v)
	try(root.Right, i+1, d, v)
}