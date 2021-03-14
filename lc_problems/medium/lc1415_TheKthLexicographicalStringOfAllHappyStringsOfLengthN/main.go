func getHappyString(n int, k int) string {
	nums := make([]int, n)
	nums[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		nums[i] = 2 * nums[i+1]
	}
	arr := []byte{'a', 'b', 'c'}
	var sb strings.Builder
	k--
	for i := 0; i < n; i++ {
		pos := k / nums[i]
		if pos >= len(arr) {
			return ""
		}
		if arr[pos] == 'a' {
			sb.WriteByte('a')
			arr = []byte{'b', 'c'}
		} else if arr[pos] == 'b' {
			sb.WriteByte('b')
			arr = []byte{'a', 'c'}
		} else {
			sb.WriteByte('c')
			arr = []byte{'a', 'b'}
		}
		k %= nums[i]
	}
	return sb.String()
}