func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	counts := make([]int, 11)
	counts[1] = 9
	count := 9
	for i := 2; i < 10; i++ {
		count *= 10
		counts[i] = i*count + counts[i-1]
	}

	k := 0
	for counts[k] < n {
		k++
	}
	n -= counts[k-1]
	num := (n - 1) / k
	base := 1
	for i := 0; i < k-1; i++ {
		base *= 10
	}
	num += base
	res := (n - 1) % k
	return int(strconv.Itoa(num)[res] - '0')
}