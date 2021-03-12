type Pair struct {
	node int
	prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	probs := make(map[int][]Pair)
	ne := len(edges)
	for i := 0; i < ne; i++ {
		edge := edges[i]
		probs[edge[0]] = append(probs[edge[0]], Pair{edge[1], succProb[i]})
		probs[edge[1]] = append(probs[edge[1]], Pair{edge[0], succProb[i]})
	}

	result := make([]float64, n)
	result[start] = 1.0

	queue := make([]Pair, 1)
	queue[0] = Pair{start, float64(1.0)}
	lq := 1
	for lq > 0 {
		for i := 0; i < lq; i++ {
			curNode := queue[i]
			nextNodes := probs[curNode.node]
			for j := 0; j < len(nextNodes); j++ {
				nextPair := nextNodes[j]
				nextNode := nextPair.node
				nextProb := nextPair.prob * curNode.prob
				if nextProb <= result[nextNode] {
					continue
				}
				result[nextNode] = nextProb
				if nextNode == end {
					continue
				}
				queue = append(queue, Pair{nextNode, nextProb})
			}

		}
		queue = queue[lq:]
		lq = len(queue)
	}

	return result[end]
}