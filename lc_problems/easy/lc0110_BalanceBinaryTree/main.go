import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func getMaxHeight(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	lh := getMaxHeight(root.Left)
	if lh == -1 {
		return -1
	}
	rh := getMaxHeight(root.Right)
	if rh == -1 {
		return -1
	}

	if lh > rh {
		if lh-rh > 1 {
			return -1
		}
		return 1 + lh
	} else {
		if rh-lh > 1 {
			return -1
		}
		return 1 + rh
	}
}

func isBalanced(root *tree.TreeNode) bool {
	if getMaxHeight(root) == -1 {
		return false
	}
	return true
}