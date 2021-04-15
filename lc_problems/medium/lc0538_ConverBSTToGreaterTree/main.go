/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root.Right != nil {
			helper(root.Right)
		}
		root.Val += sum
		sum = root.Val
		if root.Left != nil {
			helper(root.Left)
		}
	}
	if root != nil {
		helper(root)
	}

	return root
}