/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1 == nil {
		return quadTree2
	}
	if quadTree2 == nil {
		return quadTree1
	}
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return quadTree1
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return quadTree2
		}
		return quadTree1
	}

	newTree := new(Node)
	newTree.TopLeft = intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	newTree.TopRight = intersect(quadTree1.TopRight, quadTree2.TopRight)
	newTree.BottomLeft = intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	newTree.BottomRight = intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if !newTree.TopLeft.IsLeaf {
		return newTree
	}
	val := newTree.TopLeft.Val
	if !newTree.BottomLeft.IsLeaf || newTree.BottomLeft.Val != val {
		return newTree
	}
	if !newTree.TopRight.IsLeaf || newTree.TopRight.Val != val {
		return newTree
	}
	if !newTree.BottomRight.IsLeaf || newTree.BottomRight.Val != val {
		return newTree
	}
	return &Node{
		IsLeaf: true,
		Val:    val,
	}

}