type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	if n == 1 {
		return &Node{
			grid[0][0] == 1, true, nil, nil, nil, nil,
		}
	}

	var helper func(rowStart, rowEnd, colStart, colEnd int) (*Node, int)
	helper = func(rowStart, rowEnd, colStart, colEnd int) (*Node, int) {
		if rowStart == rowEnd {
			return &Node{
				grid[rowStart][colStart] == 1, true, nil, nil, nil, nil,
			}, grid[rowStart][colStart]
		}

		halfWidth := (rowEnd - rowStart + 1) >> 1
		rowMid2 := rowStart + halfWidth
		rowMid1 := rowMid2 - 1
		colMid2 := colStart + halfWidth
		colMid1 := colMid2 - 1
		topLeftNode, topLeft := helper(rowStart, rowMid1, colStart, colMid1)
		topRightNode, topRight := helper(rowStart, rowMid1, colMid2, colEnd)
		bottomLeftNode, bottomLeft := helper(rowMid2, rowEnd, colStart, colMid1)
		bottomRightNode, bottomRight := helper(rowMid2, rowEnd, colMid2, colEnd)
		if topLeft != -1 && topLeft == topRight && topLeft == bottomLeft && topLeft == bottomRight {
			return &Node{
				grid[rowStart][colStart] == 1, true, nil, nil, nil, nil,
			}, topLeft
		}
		return &Node{
			true, false, topLeftNode, topRightNode, bottomLeftNode, bottomRightNode,
		}, -1
	}

	root, v := helper(0, n-1, 0, n-1)
	if v == -1 {
		return root
	}
	return &Node{
		grid[0][0] == 1, true, nil, nil, nil, nil,
	}
}