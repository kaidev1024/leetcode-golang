
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			root.Left = buildTree(inorder[:i], postorder[:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
			break
		}
	}
	return root
}