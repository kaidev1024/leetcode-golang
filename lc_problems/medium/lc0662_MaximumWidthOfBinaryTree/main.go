type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type info struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []info{info{node: root,
		pos: 0,
	}}
	res := 1
	for len(queue) > 0 {
		l := len(queue)
		if queue[l-1].pos-queue[0].pos+1 > res {
			res = queue[l-1].pos - queue[0].pos + 1
		}
		for _, v := range queue {
			if v.node.Left != nil {
				queue = append(queue, info{v.node.Left, v.pos * 2})
			}
			if v.node.Right != nil {
				queue = append(queue, info{v.node.Right, v.pos*2 + 1})
			}
		}
		queue = queue[l:]
	}
	return res
}