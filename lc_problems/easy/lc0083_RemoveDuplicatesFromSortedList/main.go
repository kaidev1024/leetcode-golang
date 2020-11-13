package main

import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/list"
)

func deleteDuplicates(head *list.ListNode) *list.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	iter := head
	for iter != nil && iter.Next != nil {
		if iter.Val == iter.Next.Val {
			iter.Next = iter.Next.Next
		} else {
			iter = iter.Next
		}
	}
	return head
}

var l = list.NewList([]int{1, 1, 2})

func main() {
	result := deleteDuplicates(l)
	fmt.Println(result.String())
}
