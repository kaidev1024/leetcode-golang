
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)

	var helper func(root *Node)

	helper = func(root *Node) {
		arr = append(arr, root.Val)
		for _, n := range root.Children {
			helper(n)
		}

	}
	helper(root)
	return arr
}