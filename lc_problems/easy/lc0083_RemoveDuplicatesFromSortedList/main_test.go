package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	result := deleteDuplicates(l)
	if result.String() != "1->2" {
		t.Error("wrong list")
	}
}
