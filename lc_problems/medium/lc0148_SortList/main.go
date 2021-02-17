
type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := &ListNode{0, nil}
	iter := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			iter.Next = l1
			l1 = l1.Next
			iter = iter.Next
		} else {
			iter.Next = l2
			l2 = l2.Next
			iter = iter.Next
		}
	}
	if l1 != nil {
		iter.Next = l1
	}
	if l2 != nil {
		iter.Next = l2
	}
	return dummy.Next
}

func split(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	dummy := &ListNode{0, head}
	slow := dummy
	fast := dummy
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	l := slow.Next
	slow.Next = nil
	return dummy.Next, l
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	l1, l2 := split(head)
	l1 = sortList(l1)
	l2 = sortList(l2)
	return merge(l1, l2)
}