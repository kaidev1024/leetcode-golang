/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	parent, delNode, isLeft := searchNode(root, key)
	if delNode == nil {
		return root
	}

	if delNode.Right == nil {
		if parent == nil {
			return delNode.Left
		} else {
			if isLeft {
				parent.Left = delNode.Left
			} else {
				parent.Right = delNode.Left
			}
		}
	} else {
		delNode.Val = findMinRight(delNode)
	}
	return root
}

func searchNode(root *TreeNode, key int) (parent, delNode *TreeNode, isLeft bool) {
	for {
		if root == nil {
			return nil, nil, false
		}
		if root.Val == key {
			delNode = root
			return
		}
		parent = root
		if root.Val > key {
			root = root.Left
			isLeft = true
		} else {
			root = root.Right
			isLeft = false
		}
	}
	return nil, nil, false
}

func findMinRight(delNode *TreeNode) int {
	parent := delNode
	minNode := parent.Right
	isLeft := false
	for minNode.Left != nil {
		parent = minNode
		minNode = minNode.Left
		isLeft = true
	}

	if isLeft {
		parent.Left = minNode.Right
	} else {
		parent.Right = minNode.Right
	}

	return minNode.Val
}