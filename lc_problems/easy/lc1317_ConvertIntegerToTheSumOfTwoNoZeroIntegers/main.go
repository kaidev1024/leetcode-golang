func getNoZeroIntegers(n int) []int {
	n1 := 0
	n2 := 0

	mul := 1
	for n > 0 {
		fmt.Println(n, n1, n2)
		res := n % 10
		n /= 10

		if res == 0 {
			n--
			n1 += 5 * mul
			n2 += 5 * mul
		} else if res == 1 {
			if n > 0 {
				n--
				n1 += 6 * mul
				n2 += 5 * mul
			} else {
				n1 += mul
			}

		} else if res%2 == 0 {
			n1 += (res >> 1) * mul
			n2 += (res >> 1) * mul
		} else {
			half := res / 2
			n1 += half * mul
			n2 += (half + 1) * mul
		}
		mul *= 10
	}
	return []int{n1, n2}
}
