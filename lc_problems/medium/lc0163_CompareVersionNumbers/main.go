func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	n1 := len(v1)
	n2 := len(v2)

	n := n1
	if n1 > n2 {
		dif := n1 - n2
		for dif > 0 {
			v2 = append(v2, "0")
			dif--
		}
	}
	if n2 > n1 {
		dif := n2 - n1
		for dif > 0 {
			v1 = append(v1, "0")
			dif--
		}
		n = n2
	}

	for i := 0; i < n; i++ {
		a1, _ := strconv.Atoi(v1[i])
		a2, _ := strconv.Atoi(v2[i])
		if a1 < a2 {
			return -1
		}
		if a1 > a2 {
			return 1
		}
	}
	return 0
}