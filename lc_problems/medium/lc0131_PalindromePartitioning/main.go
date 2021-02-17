func partition(s string) [][]string {
	n := len(s)
	if n == 1 {
		return [][]string{[]string{s}}
	}

	matrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		matrix[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		matrix[i][i+1] = s[i] == s[i+1]
	}

	for l := 3; l <= n; l++ {
		for start := 0; start < n-l+1; start++ {
			end := start + l - 1
			if s[start] == s[end] {
				matrix[start][end] = matrix[start+1][end-1]
			}
		}
	}

	result := make([][]string, 0)
	var dfs func(start int, arr []string)
	dfs = func(start int, arr []string) {
		if start == n {
			arrCopy := make([]string, len(arr))
			copy(arrCopy, arr)
			result = append(result, arrCopy)
			return
		}
		for i := start; i < n; i++ {
			if matrix[start][i] {
				arr = append(arr, s[start:i+1])
				dfs(i+1, arr)
				arr = arr[:len(arr)-1]
			}
		}
	}
	dfs(0, []string{})
	return result
}