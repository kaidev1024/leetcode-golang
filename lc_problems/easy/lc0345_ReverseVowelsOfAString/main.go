package leetcode

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func reverseVowels(s string) string {
	arr := []byte(s)
	for l, r := 0, len(s)-1; l < r; {
		for ; l < r && !isVowel(arr[l]); l++ {
		}
		for ; l < r && !isVowel(arr[r]); r-- {
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return string(arr)
}
