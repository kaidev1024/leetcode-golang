
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			inorder(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			inorder(root.Right)
		}
	}

	inorder(root)

	result := &TreeNode{arr[0], nil, nil}
	iter := result
	for i, n := 1, len(arr); i < n; i++ {
		iter.Right = &TreeNode{arr[i], nil, nil}
		iter = iter.Right
	}
	return result
}