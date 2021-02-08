
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	level := make([]*TreeNode, 0)
	sum := 0
	result := make([]float64, 0)

	level = append(level, root)
	l := 1

	for l > 0 {
		pos := 0
		for ; pos < l; pos++ {
			tn := level[pos]
			sum += tn.Val
			if tn.Left != nil {
				level = append(level, tn.Left)
			}
			if tn.Right != nil {
				level = append(level, tn.Right)
			}
		}
		result = append(result, float64(sum)/float64(l))
		sum = 0
		level = level[pos:]
		l = len(level)
	}
	return result
}