func minSetSize(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	var count []int
	for _, val := range m {
		count = append(count, val)
	}
	sort.Ints(count)
	s := 0
	for i := len(count) - 1; i >= 0; i-- {
		s += count[i]
		if s >= len(arr)/2 {
			return len(count) - i
		}
	}
	return len(arr) / 2
}