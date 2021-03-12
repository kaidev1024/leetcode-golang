
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	result := 0
	var helper func(root *TreeNode) []int
	helper = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{1}
		}

		leftValues := helper(root.Left)
		rightValues := helper(root.Right)
		arr := make([]int, 0)
		for _, l := range leftValues {
			if l+1 < distance {
				arr = append(arr, l+1)
			}
		}
		for _, r := range rightValues {
			if r+1 < distance {
				arr = append(arr, r+1)
			}
		}

		for _, l := range leftValues {

			for _, r := range rightValues {

				sum := l + r
				if sum > distance {
					continue
				}
				result++
			}
		}
		fmt.Println(arr)
		return arr
	}
	helper(root)
	return result
}