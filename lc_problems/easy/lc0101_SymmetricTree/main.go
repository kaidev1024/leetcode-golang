import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func areSymmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && areSymmetric(t1.Left, t2.Right) && areSymmetric(t1.Right, t2.Left)
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || areSymmetric(root.Left, root.Right)
}

