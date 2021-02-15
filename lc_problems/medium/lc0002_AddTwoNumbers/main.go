type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	iter := dummy
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next

		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		iter.Next = &ListNode{sum % 10, nil}
		sum /= 10
		iter = iter.Next

	}
	if sum == 1 {
		iter.Next = &ListNode{1, nil}
	}
	return dummy.Next
}