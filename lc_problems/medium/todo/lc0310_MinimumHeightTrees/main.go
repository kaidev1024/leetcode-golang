func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	leaves := make([]int, 0, n)
	for k, v := range graph {
		if len(v) == 1 {
			leaves = append(leaves, k)
		}
	}

	for n > 2 {
		n -= len(leaves)
		newLeaves := []int{}
		for _, node := range leaves {

			target := graph[node][0]
			l := len(graph[target])
			for i := 0; i < l; i++ {
				if graph[target][i] == node {
					graph[target][i] = graph[target][l-1]
					break
				}
			}
			graph[target] = graph[target][:l-1]
			if l == 2 {
				newLeaves = append(newLeaves, target)
			}
		}
		leaves = newLeaves

	}
	return leaves
}