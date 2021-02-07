func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	arr := make([]byte, 9)

	pos := 8
	for num > 0 {
		res := num % 7
		arr[pos] = '0' + byte(res)
		pos--
		num /= 7
	}
	return sign + string(arr[pos+1:])
}