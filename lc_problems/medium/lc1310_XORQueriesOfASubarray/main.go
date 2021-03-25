func xorQueries(arr []int, queries [][]int) []int {

	result := make([]int, len(queries))
	n := len(arr)
	if n == 0 || len(queries) == 0 {
		return result
	}

	xor := make([]int, n+1)

	for i := 1; i <= len(arr); i++ {
		xor[i] = arr[i-1] ^ xor[i-1]
	}

	for idx, query := range queries {
		result[idx] = xor[query[1]+1] ^ xor[query[0]]
	}

	return result
}