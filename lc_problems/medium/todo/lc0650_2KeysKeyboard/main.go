func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	f := isPrime(n)
	if f == 0 {
		return n
	}
	return minSteps(n/f) + f
}

func isPrime(n int) int {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return i
		}
	}
	return 0
}