/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, nxt := head, head
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l1.Val == 0 {
		return l2
	}
	if l2 == nil || l2.Val == 0 {
		return l1
	}

	l1 = reverse(l1)
	l2 = reverse(l2)
	carry := 0

	dummy := ListNode{0, nil}
	cur := &dummy

	for l1 != nil || l2 != nil || carry != 0 {
		l1v := 0
		l2v := 0
		if l1 != nil {
			l1v = l1.Val
		}
		if l2 != nil {
			l2v = l2.Val
		}

		ret := l1v + l2v + carry
		val := ret % 10
		carry = ret / 10

		cur.Next = &ListNode{val, nil}
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return reverse(dummy.Next)
}