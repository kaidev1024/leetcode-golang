package structs

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(v int) *ListNode {
	return &ListNode{
		Val: v,
	}
}

func (ln *ListNode) Append(next *ListNode) {
	ln.Next = next
}

func NewList(arr []int) *ListNode {
	head := NewListNode(0)
	iterator := head
	for _, val := range arr {
		next := NewListNode(val)
		iterator.Next = next
		iterator = next
	}
	return head.Next
}
