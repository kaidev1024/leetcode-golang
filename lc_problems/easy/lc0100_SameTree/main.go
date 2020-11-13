package main

import (
	"fmt"

	"github.com/kaidev1024/leetcode-golang/structs/tree"
)

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

var t1 = tree.NewTreeNode(1)
var t2 = tree.NewTreeNode(1)
var t3 = tree.NewTreeNode(3)

func main() {
	fmt.Println(isSameTree(t1, t2))
	fmt.Println(isSameTree(t1, t3))
}
