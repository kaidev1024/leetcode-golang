package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(ln *ListNode) int {
	n := 0
	for ln != nil {
		n++
		ln = ln.Next
	}
	return n
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	na := listLen(headA)
	nb := listLen(headB)
	if na == 0 || nb == 0 {
		return nil
	}

	iterA := headA
	iterB := headB
	dif := na - nb
	if na < nb {
		iterA = headB
		iterB = headA
		dif = -dif
	}
	for dif > 0 {
		iterA = iterA.Next
		dif--
	}

	for {
		if iterA == iterB {
			return iterA
		}
		iterA = iterA.Next
		iterB = iterB.Next
	}
}
