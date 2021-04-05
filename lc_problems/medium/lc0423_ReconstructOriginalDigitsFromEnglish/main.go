func originalDigits(s string) string {
	dict := make([]int, 26)
	for i := 0; i < len(s); i++ {
		dict[int(s[i]-'a')]++
	}

	res := make([]string, 10)
	res[0] = helper('z', dict, "zero", "0")
	res[6] = helper('x', dict, "six", "6")
	res[2] = helper('w', dict, "two", "2")
	res[4] = helper('u', dict, "four", "4")
	res[5] = helper('f', dict, "five", "5")
	res[1] = helper('o', dict, "one", "1")
	res[7] = helper('s', dict, "seven", "7")
	res[3] = helper('r', dict, "three", "3")
	res[8] = helper('t', dict, "eight", "8")
	res[9] = helper('i', dict, "nine", "9")

	return strings.Join(res, "")
}

func helper(b byte, dict []int, s string, num string) string {
	v := dict[int(b-'a')]
	for i := 0; i < len(s); i++ {
		dict[int(s[i]-'a')] -= v
	}
	return strings.Repeat(num, v)
}