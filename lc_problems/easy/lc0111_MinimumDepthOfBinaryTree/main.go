import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	result := 1

	for {
		nq := len(queue)
		if nq == 0 {
			break
		}

		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return result
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result++
		queue = queue[nq:]
	}
	return result
}