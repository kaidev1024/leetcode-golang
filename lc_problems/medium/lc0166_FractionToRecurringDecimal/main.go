func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func fractionToDecimal(numerator int, denominator int) string {

	n := numerator / denominator
	res := numerator % denominator
	nStr := strconv.Itoa(n)
	if res == 0 {
		return nStr
	}
	sign := ""
	if numerator > 0 && denominator < 0 {
		sign = "-"
	}
	if numerator < 0 && denominator > 0 {
		sign = "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	n = numerator / denominator
	nStr = strconv.Itoa(n)
	res = numerator % denominator
	m := make(map[int]int)
	arr := make([]int, 0)
	pos := 0
	for res > 0 {
		if _, ok := m[res]; ok {
			break
		}
		m[res] = pos
		pos++
		res *= 10
		arr = append(arr, res/denominator)
		res %= denominator
	}
	if res == 0 {
		sb := new(strings.Builder)
		for _, v := range arr {
			sb.WriteByte(byte(v) + '0')
		}
		return sign + nStr + "." + sb.String()
	}
	ind, _ := m[res]
	sb := new(strings.Builder)
	for i := 0; i < ind; i++ {
		sb.WriteByte('0' + byte(arr[i]))
	}
	sb.WriteByte('(')
	for i := ind; i < len(arr); i++ {
		sb.WriteByte('0' + byte(arr[i]))
	}
	sb.WriteByte(')')
	return sign + nStr + "." + sb.String()
}