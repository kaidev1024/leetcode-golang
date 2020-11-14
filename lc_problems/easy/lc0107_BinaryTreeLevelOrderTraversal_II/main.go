import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func levelOrderBottom(root *tree.TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*tree.TreeNode{root}
	for {
		l := len(queue)
		if l == 0 {
			return result
		}
		arr := make([]int, 0, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{arr}, result...)
		queue = queue[l:]
	}
}