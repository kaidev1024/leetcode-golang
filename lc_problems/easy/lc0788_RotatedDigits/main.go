func rotatedDigits(N int) int {
	var cnt int
	for i := 2; i <= N; i++ {
		if isValid(i) {
			cnt++
		}
	}
	return cnt
}

func isValid(num int) bool {
	var ret bool
	for ; num > 0; num /= 10 {
		switch num % 10 {
		case 2, 5, 6, 9:
			ret = true
		case 3, 4, 7:
			return false
		}
	}
	return ret
}