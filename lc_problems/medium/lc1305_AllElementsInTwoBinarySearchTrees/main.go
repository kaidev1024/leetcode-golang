type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(t *TreeNode, list *[]int) {
	if t == nil {
		return
	}
	DFS(t.Left, list)
	*list = append(*list, t.Val)
	DFS(t.Right, list)
}
func merge(list1, list2 []int) []int {
	var res []int
	p1 := 0
	p2 := 0
	for p1 < len(list1) || p2 < len(list2) {
		if p2 >= len(list2) || p1 < len(list1) && list1[p1] <= list2[p2] {
			res = append(res, list1[p1])
			p1++
		} else if p1 >= len(list1) || p2 < len(list2) && list1[p1] > list2[p2] {
			res = append(res, list2[p2])
			p2++
		}
	}
	return res
}
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var list1, list2 []int
	DFS(root1, &list1)
	DFS(root2, &list2)
	return merge(list1, list2)
}