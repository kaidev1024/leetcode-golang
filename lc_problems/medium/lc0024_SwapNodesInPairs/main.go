
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	iter := dummy
	for {
		first := iter.Next
		if first == nil || first.Next == nil {
			break
		}
		second := first.Next
		first.Val, second.Val = second.Val, first.Val
		iter = second
	}
	return dummy.Next
}