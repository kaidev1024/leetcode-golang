
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	arr := make([]int, 0)

	var helper func(root *Node)

	helper = func(root *Node) {

		for _, n := range root.Children {
			helper(n)
		}
		arr = append(arr, root.Val)
	}
	helper(root)
	return arr
}