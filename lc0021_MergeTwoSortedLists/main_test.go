package main

import (
	"testing"

	list "github.com/kaidev1024/leetcode-golang/structs/list"
)

func TestMergeTwoLists(t *testing.T) {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := list.NewList(arr1)
	l2 := list.NewList(arr2)
	result := mergeTwoLists(l1, l2)

	expectResult := "1->1->2->3->4->4"
	if result.String() != expectResult {
		t.Error("wrong list")
	}
}
