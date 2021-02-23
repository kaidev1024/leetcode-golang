type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	nodes := []*TreeNode{root}
	for {
		n := len(nodes)
		if n == 0 {
			break
		}
		cur := nodes[n-1]
		result = append([]int{cur.Val}, result...)
		nodes = nodes[:n-1]
		if cur.Left != nil {
			nodes = append(nodes, cur.Left)
		}
		if cur.Right != nil {
			nodes = append(nodes, cur.Right)
		}

	}
	return result
}