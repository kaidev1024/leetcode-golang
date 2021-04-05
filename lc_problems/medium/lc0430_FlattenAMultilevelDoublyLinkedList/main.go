type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func helper(root *Node) *Node {
	iter := root
	for iter != nil {
		next := iter.Next
		if iter.Child == nil {
			if next == nil {
				break
			}
			iter = next
			continue
		}

		childHead := iter.Child
		iter.Child = nil
		childTail := helper(childHead)
		childHead.Prev = iter
		iter.Next = childHead
		childTail.Next = next
		if next != nil {
			next.Prev = childTail
			iter = next
		} else {
			return childTail
		}
	}
	return iter
}

func flatten(root *Node) *Node {
	helper(root)
	return root
}