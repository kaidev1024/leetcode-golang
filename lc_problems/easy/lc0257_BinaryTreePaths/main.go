package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)

	var helper func(root *TreeNode, s string)

	helper = func(root *TreeNode, s string) {
		if root.Left == nil && root.Right == nil {
			result = append(result, s)
			return
		}

		if root.Left != nil {
			helper(root.Left, s+"->"+strconv.Itoa(root.Left.Val))
		}
		if root.Right != nil {
			helper(root.Right, s+"->"+strconv.Itoa(root.Right.Val))
		}

	}

	helper(root, strconv.Itoa(root.Val))
	return result
}
