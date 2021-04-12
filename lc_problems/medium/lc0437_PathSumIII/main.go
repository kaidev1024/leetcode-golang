type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	result := 0
	var helper func(root *TreeNode, cur int, counts map[int]int)
	helper = func(root *TreeNode, cur int, counts map[int]int) {
		cur += root.Val
		if v, ok := counts[cur-sum]; ok {
			result += v
		}
		counts[cur]++
		if root.Left != nil {
			helper(root.Left, cur, counts)
		}
		if root.Right != nil {
			helper(root.Right, cur, counts)
		}
		counts[cur]--
	}
	counts := make(map[int]int)
	counts[0] = 1
	helper(root, 0, counts)
	return result
}