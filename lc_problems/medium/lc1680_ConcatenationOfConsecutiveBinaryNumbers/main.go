func convert(n int) int {
	l := 0
	for n > 0 {
		l++
		n /= 2
	}
	return l
}

func concatenatedBinary(n int) int {
	s := 0

	for i := 1; i <= n; i++ {
		l := convert(i)
		for l > 0 {
			s *= 2
			l--
		}
		s += i
		s %= 1000000007
	}
	return s
}