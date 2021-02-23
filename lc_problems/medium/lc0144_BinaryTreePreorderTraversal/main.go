
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
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
		result = append(result, cur.Val)
		nodes = nodes[:n-1]
		if cur.Right != nil {
			nodes = append(nodes, cur.Right)
		}
		if cur.Left != nil {
			nodes = append(nodes, cur.Left)
		}
	}
	return result
}