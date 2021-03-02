func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indegree := make([]int, n)
	res := make([]int, 0)

	for _, edge := range edges {
		indegree[edge[1]]++
	}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}