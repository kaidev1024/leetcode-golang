package main

import (
	"fmt"

	list "github.com/kaidev1024/leetcode-golang/structs/list"
)

func mergeTwoLists(l1 *list.ListNode, l2 *list.ListNode) *list.ListNode {
	var head list.ListNode
	iterator := &head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			iterator.Next = l1
			iterator = iterator.Next
			l1 = l1.Next
		} else {
			iterator.Next = l2
			iterator = iterator.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		iterator.Next = l1
	}
	if l2 != nil {
		iterator.Next = l2
	}
	return head.Next
}

func main() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := list.NewList(arr1)
	l2 := list.NewList(arr2)
	result := mergeTwoLists(l1, l2)
	fmt.Println(result)
}
