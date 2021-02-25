
type Node struct {
	Val       int
	Neighbors []*Node
}

var m = make(map[*Node]*Node)

func NewNode(node *Node) *Node {
	l := len(node.Neighbors)

	neighbors := make([]*Node, l)
	return &Node{
		Val:       node.Val,
		Neighbors: neighbors,
	}
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := m[node]; ok {
		return n
	}

	newNode := NewNode(node)
	m[node] = newNode
	for i, v := range node.Neighbors {
		newNode.Neighbors[i] = cloneGraph(v)
	}

	return newNode
}