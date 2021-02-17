
type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}
	return &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
}