func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]int, n)
	nums[0] = 1

	np := len(primes)
	pos := make([]int, np)

	for i := 1; i < n; i++ {
		v := primes[0] * nums[pos[0]]
		for j := 1; j < np; j++ {
			curV := primes[j] * nums[pos[j]]
			if curV < v {
				v = curV
			}
		}
		for j := 0; j < np; j++ {
			if v == primes[j]*nums[pos[j]] {
				pos[j]++
			}
		}
		nums[i] = v
	}
	return nums[n-1]
}