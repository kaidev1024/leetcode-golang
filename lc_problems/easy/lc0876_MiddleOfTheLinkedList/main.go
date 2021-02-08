
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy

	for {
		if fast.Next != nil {
			fast = fast.Next
		}
		if fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
		if fast.Next == nil {
			break
		}
	}
	return slow.Next
}