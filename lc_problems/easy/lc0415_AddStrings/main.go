package leetcode

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	if l1 == 0 {
		return num2
	}
	l2 := len(num2)
	if l2 == 0 {
		return num1
	}
	if l1 > l2 {
		return addStrings(num2, num1)
	}

	bytes1 := []byte(num1)
	bytes2 := []byte(num2)
	result := make([]byte, l2+1)
	l := len(result)
	var carry byte = 0
	for i := 0; i < l2; i++ {
		sum := carry + bytes2[l2-1-i] - '0'
		if l1-1-i >= 0 {
			sum += bytes1[l1-1-i] - '0'
		}
		carry = sum / 10
		result[l-1-i] = sum%10 + '0'
	}
	if carry == 1 {
		result[0] = 1 + '0'
	} else {
		result = result[1:]
	}
	return string(result)
}
