type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		max := 0
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i == 0 {
				max = node.Val
			} else {
				if node.Val > max {
					max = node.Val
				}
			}
		}
		ans = append(ans, max)
	}

	return ans
}