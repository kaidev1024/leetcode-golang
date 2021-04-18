func judgeSquareSum(c int) bool {
	for i := 0; 2*i*i <= c; i++ {
		res := c - i*i
		temp := int(math.Sqrt(float64(res)))
		if res == temp*temp {
			return true
		}
	}
	return false
}