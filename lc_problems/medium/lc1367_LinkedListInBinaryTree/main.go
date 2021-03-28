type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var helper func(head *ListNode, root *TreeNode) bool
	helper = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}
		if root == nil || head.Val != root.Val {
			return false
		}
		return helper(head.Next, root.Left) || helper(head.Next, root.Right)
	}

	if root == nil {
		return false
	}

	if helper(head, root) {
		return true
	}

	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}