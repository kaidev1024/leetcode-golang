func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]int)
	for i, val := range arr {
		m[val] = i
	}

	for _, piece := range pieces {
		ind := m[piece[0]]
		for _, v := range piece {
			if ind == len(arr) {
				return false
			}
			if v != arr[ind] {
				return false
			}
			ind++
		}
	}
	return true
}