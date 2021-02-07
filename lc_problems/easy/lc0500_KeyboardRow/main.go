func findWords(words []string) []string {
	dict := make(map[byte]int)
	strs := []string{"QWERTYUIOPqwertyuiop", "ASDFGHJKLasdfghjkl", "ZXCVBNMzxcvbnm"}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(strs[i]); j++ {
			dict[strs[i][j]] = i + 1
		}
	}

	var res []string
LOOP:
	for _, s := range words {
		tar := dict[s[0]]
		for i := 1; i < len(s); i++ {
			if dict[s[i]] != tar {
				continue LOOP
			}
		}
		res = append(res, s)
	}
	return res
}