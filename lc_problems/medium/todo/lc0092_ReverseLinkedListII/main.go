
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	shadowHead := &ListNode{Next: head}
	prevNode := shadowHead
	for i := 1; i < m; i++ {
		prevNode = prevNode.Next
		head = head.Next
	}
	mNode := head
	nNode := head.Next
	for i := m; i < n; i++ {
		nextNNode := nNode.Next
		nNode.Next = head
		head = nNode
		nNode = nextNNode
	}
	mNode.Next = nNode
	prevNode.Next = head
	return shadowHead.Next
}