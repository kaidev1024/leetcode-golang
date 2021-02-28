
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	evenHead := head.Next
	evenIter := evenHead
	for evenIter != nil && evenIter.Next != nil {
		temp := evenIter.Next
		evenIter.Next = temp.Next
		temp.Next = evenHead
		odd.Next = temp
		odd = temp
		evenIter = evenIter.Next
	}
	return head
}