func findJudge(N int, trust [][]int) int {
	indegree := make([]int, N+1)
	outdegree := make([]int, N+1)
	for _, pair := range trust {
		indegree[pair[1]]++
		outdegree[pair[0]]++
	}
	judge := -1
	for i := 1; i <= N; i++ {
		if indegree[i] == N-1 && outdegree[i] == 0 {
			if judge == -1 {
				judge = i
			} else {
				return -1
			}
		}
	}
	return judge
}