package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	next := head
	var prev, cur *ListNode
	for next != nil {
		cur = next
		next = cur.Next
		cur.Next = prev
		prev = cur
	}
	return cur
}

// recursive
// func reverseList(head *ListNode) *ListNode {
//     if head == nil || head.Next == nil {
//         return head
//     }

//     next := head.Next

//     result := reverseList(next)
//     next.Next = head
//     head.Next = nil
//     return result
// }
