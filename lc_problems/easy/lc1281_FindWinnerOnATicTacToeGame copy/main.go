func subtractProductAndSum(n int) int {
	p := 1
	s := 0
	for n > 0 {
		res := n % 10
		n /= 10
		p *= res
		s += res
	}
	return p - s
}