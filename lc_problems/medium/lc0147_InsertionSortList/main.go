
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	prev := head
	cur := head.Next

	for cur != nil {
		if cur.Val > prev.Val {
			prev = cur
			cur = cur.Next
		} else {
			iter := dummy
			for iter.Next.Val < cur.Val {
				iter = iter.Next
			}
			prev.Next = cur.Next
			cur.Next = iter.Next
			iter.Next = cur
			cur = prev.Next
		}
	}
	return dummy.Next
}