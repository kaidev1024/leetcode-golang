package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{0, head}
	iter := dummy
	for {
		if iter.Next == nil {
			return dummy.Next
		}
		if iter.Next.Val == val {
			iter.Next = iter.Next.Next
		} else {
			iter = iter.Next
		}
	}
}
