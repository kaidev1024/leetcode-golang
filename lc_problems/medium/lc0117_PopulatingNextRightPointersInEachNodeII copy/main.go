
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func connect(root *Node) *Node {
	var q []*Node

	if root != nil {
		q = append(q, root)

		for len(q) > 0 {
			count := len(q)
			var prev *Node
			for count > 0 {
				q[0].Next, prev = prev, q[0]
				if q[0].Right != nil {
					q = append(q, q[0].Right)
				}
				if q[0].Left != nil {
					q = append(q, q[0].Left)
				}

				q = q[1:]
				count--
			}
		}
	}
	return root
}