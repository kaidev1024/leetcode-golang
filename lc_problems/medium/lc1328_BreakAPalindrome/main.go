func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n < 2 {
		return ""
	}

	index := 0
	half := (n - 1) >> 1
	for ; index < half; index++ {
		if palindrome[index] != 'a' {
			break
		}
	}
	arr := []byte(palindrome)
	if n&1 == 0 {
		if arr[index] == 'a' {
			arr[index+1] = 'b'
		} else {
			arr[index] = 'a'
		}

		return string(arr)
	}

	if index == half {
		arr[n-1] = 'b'
	} else {
		arr[index] = 'a'
	}
	return string(arr)
}