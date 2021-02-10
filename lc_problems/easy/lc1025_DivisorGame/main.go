func divisorGame(N int) bool {
	if N == 1 {
		return false
	}
	result := make([]bool, N+1)
	result[2] = true
	for i := 3; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !result[i-j] {
				result[i] = true
				break
			}
		}
	}
	return result[N]
}