type Solution struct {
	node *ListNode
	n    int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	n := 0
	iter := head
	for iter != nil {
		n++
		iter = iter.Next
	}
	return Solution{node: head, n: n}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	k := rand.Intn(this.n)
	iter := this.node
	for k > 0 {
		iter = iter.Next
		k--
	}
	return iter.Val
}