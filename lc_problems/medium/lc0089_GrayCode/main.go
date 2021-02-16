func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	arr := []int{0, 1}
	for i := 1; i < n; i++ {
		v := 1 << i
		l := len(arr)
		for j := l - 1; j >= 0; j-- {
			arr = append(arr, v|arr[j])
		}
	}
	return arr
}