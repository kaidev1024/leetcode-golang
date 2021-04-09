func validIPAddress(IP string) string {
	str1 := strings.Split(IP, ".")
	str2 := strings.Split(IP, ":")
	if len(str1) != 4 && len(str2) != 8 {
		return "Neither"
	}

	if len(str1) == 4 {
		for _, v := range str1 {
			n, err := strconv.Atoi(v)
			if err != nil || n < 0 || n > 255 {
				goto ipv6
			}
			if n != 0 && v[0] == '0' {
				goto ipv6
			}
			if n == 0 && len(v) > 1 {
				goto ipv6
			}
		}
		return "IPv4"
	}
ipv6:
	if len(str2) == 8 {
		for _, v := range str2 {
			if len(v) == 0 || len(v) > 4 {
				return "Neither"
			}
			for i := 0; i < len(v); i++ {
				if (v[i] >= '0' && v[i] <= '9') ||
					(v[i] >= 'a' && v[i] <= 'f') ||
					(v[i] >= 'A' && v[i] <= 'F') {
					continue
				} else {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}