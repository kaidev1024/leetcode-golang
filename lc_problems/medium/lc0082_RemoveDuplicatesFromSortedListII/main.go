
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	start := dummy
	end := start.Next
	for end != nil {
		for end.Next != nil && end.Val == end.Next.Val {
			end = end.Next
		}
		if start.Next != end {
			start.Next = end.Next
			end = end.Next
		} else {
			start = end
			end = end.Next
		}
	}
	return dummy.Next
}