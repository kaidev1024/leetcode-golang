func maximum69Number(num int) int {
	arr := []byte(strconv.Itoa(num))
	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] == '6' {
			arr[i] = '9'
			break
		}
	}
	result, _ := strconv.Atoi(string(arr))
	return result
}