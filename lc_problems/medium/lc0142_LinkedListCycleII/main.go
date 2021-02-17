
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		slow = slow.Next
		if fast == nil {
			return nil
		}
		if fast == slow {
			break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}