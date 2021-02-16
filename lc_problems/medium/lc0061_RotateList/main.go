
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	// special no length case
	if head == nil {
		return head
	}

	// find length and last node
	n := head

	length := 1
	for n.Next != nil {
		n = n.Next
		length++
	}
	last := n

	// remove loops from k
	k = k % length

	// no rotation
	if k == 0 {
		return head
	}

	// point last to head
	last.Next = head

	// find new last
	n = head
	for i := 1; i < length-k; i++ {
		n = n.Next
	}

	// set next to new head
	head = n.Next
	// make n the last node
	n.Next = nil

	return head
}