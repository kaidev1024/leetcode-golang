
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	current := []*TreeNode{root}
	next := []*TreeNode{}
	for {
		level := []int{}
		for _, c := range current {
			if c == nil {
				continue
			}
			level = append(level, c.Val)
			next = append(next, c.Left, c.Right)
		}
		if len(next) == 0 {
			break
		}
		levels = append(levels, level)
		current = next
		next = nil
	}
	return levels
}