func countSubTrees(n int, edges [][]int, labels string) []int {
	m := make(map[int][]int)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}

	result := make([]int, n)
	visited := make(map[int]struct{})

	var help func(node int, positions map[byte][]int)
	help = func(node int, positions map[byte][]int) {
		visited[node] = struct{}{}
		char := labels[node]
		arr := positions[char]
		newArr := append(arr, node)
		for _, v := range newArr {
			result[v]++
		}
		positions[char] = newArr

		children := m[node]
		for _, child := range children {
			if _, ok := visited[child]; !ok {
				help(child, positions)
			}

		}
		positions[char] = arr
	}

	positions := make(map[byte][]int)

	help(0, positions)
	return result
}