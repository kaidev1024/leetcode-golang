func countPrimeSetBits(L int, R int) int {
	var cnt int
	const v = 0b010100000100010100010100010101100
	for i := L; i <= R; i++ {
		cnt += (v >> bits.OnesCount(uint(i))) & 1
	}
	return cnt
}