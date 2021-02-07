
type Node struct {
	Val      int
	Children []*Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	result := 0
	for i, n := 0, len(root.Children); i < n; i++ {
		result = max(result, maxDepth(root.Children[i]))
	}
	return 1 + result
}