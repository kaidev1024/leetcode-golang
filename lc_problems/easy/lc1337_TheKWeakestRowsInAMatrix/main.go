func kWeakestRows(mat [][]int, k int) []int {
	search := func(arr []int) int {
		l := len(arr)
		if arr[l-1] == 1 {
			return l
		}
		left := 0
		right := l - 1
		for left < right {
			mid := (right-left)/2 + left
			if arr[mid] == 1 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	type pair struct {
		index, count int
	}

	n := len(mat)
	result := make([]pair, 0, n)
	for i := 0; i < n; i++ {
		count := search(mat[i])
		result = append(result, pair{i, count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count == result[j].count {
			return result[i].index < result[j].index
		}
		return result[i].count < result[j].count
	})

	r := make([]int, k)
	for i := 0; i < k; i++ {
		r[i] = result[i].index
	}
	return r
}