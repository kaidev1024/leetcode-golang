package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(l *ListNode) *ListNode {
	var prev, next *ListNode
	curr := l

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func split(l *ListNode) (*ListNode, *ListNode) {
	if l == nil {
		return nil, nil
	}
	head := &ListNode{
		Val:  0,
		Next: l,
	}
	slow := head
	fast := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}
	secondHalf := slow.Next
	slow.Next = nil
	return l, secondHalf
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	firstHalf, secondHalf := split(head)
	secondHalfReverse := reverse(secondHalf)
	for secondHalfReverse != nil {
		if firstHalf.Val != secondHalfReverse.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalfReverse = secondHalfReverse.Next
	}
	return true
}
