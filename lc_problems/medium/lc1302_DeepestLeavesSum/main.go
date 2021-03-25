type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	summary := 0
	currentLevel := []*TreeNode{root}
	for len(currentLevel) != 0 {
		nextLevel := []*TreeNode{}
		summary = 0
		for _, v := range currentLevel {
			summary += v.Val
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		currentLevel = nextLevel
	}
	return summary
}