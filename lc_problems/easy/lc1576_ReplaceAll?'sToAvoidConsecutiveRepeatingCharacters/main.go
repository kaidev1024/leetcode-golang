func modifyString(s string) string {
	arr := []byte(s)

	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] == '?' {
			var pos byte
			if i != 0 && i != n-1 {
				for 'a'+pos == arr[i-1] || 'a'+pos == arr[i+1] {
					pos++
				}
			} else if i == 0 && i == n-1 {

			} else if i == 0 {
				if 'a'+pos == arr[i+1] {
					pos++
				}
			} else if i == n-1 {
				if 'a'+pos == arr[i-1] {
					pos++
				}
			}
			arr[i] = 'a' + pos
		}
	}

	return string(arr)
}