
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodes []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nodes := make([]*TreeNode, 0)
	for root != nil {
		nodes = append(nodes, root)
		root = root.Left
	}
	return BSTIterator{nodes}
}

func (this *BSTIterator) Next() int {
	n := len(this.nodes)
	node := this.nodes[n-1]
	result := node.Val
	this.nodes = this.nodes[:n-1]
	node = node.Right
	for node != nil {
		this.nodes = append(this.nodes, node)
		node = node.Left
	}
	return result
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > 0
}

