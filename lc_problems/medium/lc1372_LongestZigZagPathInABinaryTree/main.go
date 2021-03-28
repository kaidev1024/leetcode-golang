type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	ret := 0

	var getLength func(root *TreeNode, n int, isLeft bool)
	getLength = func(root *TreeNode, n int, isLeft bool) {
		if n > ret {
			ret = n
		}

		if isLeft && root.Right != nil {
			getLength(root.Right, n+1, false)
		}
		if isLeft && root.Left != nil {
			getLength(root.Left, 1, true)
		}
		if !isLeft && root.Left != nil {
			getLength(root.Left, n+1, true)
		}
		if !isLeft && root.Right != nil {
			getLength(root.Right, 1, false)
		}
	}
	if root.Left != nil {
		getLength(root.Left, 1, true)
	}
	if root.Right != nil {
		getLength(root.Right, 1, false)
	}
	return ret
}
