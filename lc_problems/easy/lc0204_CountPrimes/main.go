package leetcode

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	arr := make([]bool, n)

	for i := 2; i*i < n; i++ {
		for j := i * i; j < n; j += i {
			if arr[j] {
				continue
			}
			arr[j] = true
		}
	}

	result := 0
	for i := 2; i < n; i++ {
		if !arr[i] {
			result++
		}
	}
	return result
}
