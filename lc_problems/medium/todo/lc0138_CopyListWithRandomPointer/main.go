
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dummyOld := &Node{Val: 0, Next: head}
	iterOld := dummyOld.Next

	for iterOld != nil {
		next := iterOld.Next
		iterOld.Next = &Node{Val: iterOld.Val, Next: next}
		iterOld = next
	}

	result := dummyOld.Next.Next

	iter := dummyOld.Next
	for iter != nil && iter.Next != nil {
		if iter.Random == nil {
			iter.Next.Random = nil
		} else {
			iter.Next.Random = iter.Random.Next
		}

		iter = iter.Next.Next
	}

	iter = dummyOld.Next

	for iter != nil && iter.Next != nil {
		next := iter.Next.Next
		if next == nil {
			iter.Next.Next = nil
			iter.Next = nil
			break
		}
		iter.Next.Next = next.Next
		iter.Next = next
		iter = next
	}

	return result
}