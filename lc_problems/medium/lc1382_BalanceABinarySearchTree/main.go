type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root.Left != nil {
			inorder(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			inorder(root.Right)
		}

	}
	inorder(root)
	n := len(arr)

	var construct func(l, r int) *TreeNode
	construct = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		mid := (l + r) / 2
		newNode := &TreeNode{Val: arr[mid]}
		newNode.Left = construct(l, mid-1)
		newNode.Right = construct(mid+1, r)
		return newNode
	}
	return construct(0, n-1)
}