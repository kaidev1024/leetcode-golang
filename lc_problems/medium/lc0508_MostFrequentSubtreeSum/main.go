type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	counts := make(map[int]int)
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftSum := helper(root.Left)
		rightSum := helper(root.Right)
		sum := root.Val + leftSum + rightSum
		counts[sum]++
		return sum
	}
	helper(root)

	count := 0
	result := []int{}
	for key, value := range counts {
		if value < count {
			continue
		}
		if value > count {
			count = value
			result = []int{key}
			continue
		}
		result = append(result, key)
	}

	return result
}