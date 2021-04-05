type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*Node{root}
	result := make([][]int, 0)

	n := 1
	for n > 0 {
		level := make([]int, n)
		for i := 0; i < n; i++ {
			curNode := nodes[i]
			children := curNode.Children
			level[i] = curNode.Val
			for _, child := range children {
				nodes = append(nodes, child)
			}
		}
		result = append(result, level)
		nodes = nodes[n:]
		n = len(nodes)
	}
	return result
}