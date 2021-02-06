func sumZero(n int) []int {
	result := make([]int, 0, n)
	half := n / 2
	if n%2 == 1 {
		result = append(result, 0)
	}

	for i := 1; i <= half; i++ {
		result = append(result, i, -i)
	}

	return result
}