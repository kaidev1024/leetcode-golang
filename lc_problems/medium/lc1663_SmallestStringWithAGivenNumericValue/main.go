func genStr(c byte, n int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = c
	}
	return string(arr)
}

func helper(na, nz, n int) string {
	if nz+na > n {
		la := n - nz - 1
		index := na - la
		return genStr('a', la) + genStr('a'+byte(index-1), 1) + genStr('z', nz)
	}
	return genStr('a', na) + genStr('z', nz)
}

func getSmallestString(n int, k int) string {
	nz := k / 26
	na := k % 26
	if na+nz >= n {
		return helper(na, nz, n)
	}

	dif := n - nz - na
	nzz := dif / 25
	naa := nzz * 26
	nz -= nzz
	na += naa
	if nz+na < n {
		na += 26
		nz--
	}
	return helper(na, nz, n)
}
