
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{0, head}
	slow := dummy
	for {
		for slow.Next != nil && slow.Next.Val < x {
			slow = slow.Next
		}
		fast := slow
		for fast.Next != nil && fast.Next.Val >= x {
			fast = fast.Next
		}
		if fast.Next == nil {
			break
		}
		temp := fast.Next
		fast.Next = temp.Next
		temp.Next = slow.Next
		slow.Next = temp
		slow = slow.Next
	}
	return dummy.Next
}