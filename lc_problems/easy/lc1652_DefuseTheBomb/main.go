func decrypt(code []int, k int) []int {
	l := len(code)

	result := make([]int, l)
	if k == 0 {
		return result
	}

	if k > 0 {
		sum := 0
		for i := 1; i <= k; i++ {
			sum += code[i]
		}
		for i := 0; i < l; i++ {
			result[i] = sum
			sum += code[(k+i+1+l)%l] - code[(i+1+l)%l]
		}
	}
	if k < 0 {
		k = -k
		sum := 0
		for i := 1; i <= k; i++ {
			sum += code[l-1-i]
		}
		for i := l - 1; i >= 0; i-- {
			result[i] = sum
			sum += code[(i-k-1+l)%l] - code[(i-1+l)%l]
		}
	}
	return result
}