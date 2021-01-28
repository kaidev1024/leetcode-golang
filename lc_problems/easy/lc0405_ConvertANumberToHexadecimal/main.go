package leetcode

func toHex(num int) string {
	result := make([]byte, 8)

	mask := 15

	for i := 7; i >= 0; i-- {
		v := mask & num
		if v < 10 {
			result[i] = byte(v) + '0'
		} else {
			result[i] = byte(v-10) + 'a'
		}
		num >>= 4
		if num == 0 {
			result = result[i:]

			break
		}
	}
	return string(result)
}
