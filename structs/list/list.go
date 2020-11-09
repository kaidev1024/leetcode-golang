package structs

import (
	"fmt"
	"strconv"
	"strings"
)

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

func (ln *ListNode) Print() {
	if ln == nil {
		fmt.Println("empty list")
	}
	var sb strings.Builder
	for ln.Next != nil {
		sb.WriteString(strconv.Itoa(ln.Val))
		sb.WriteString("->")
		ln = ln.Next
	}
	sb.WriteString(strconv.Itoa(ln.Val))
	sb.WriteString("\n")
	fmt.Println(sb.String())
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
