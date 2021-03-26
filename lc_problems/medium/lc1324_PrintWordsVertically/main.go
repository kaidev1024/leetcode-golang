func printVertically(s string) []string {
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return nil
	}
	//  find the max length of the words to difine the size of the two-dimensional array
	maxL := len(words[0])
	for i := 1; i < len(words); i++ {
		if maxL < len(words[i]) {
			maxL = len(words[i])
		}
	}
	res := make([]string, maxL)
	// fill in the words vertically
	for i := 0; i < len(words); i++ {
		for j := 0; j < maxL; j++ {
			if j < len(words[i]) {
				res[j] = res[j] + string(words[i][j])
			} else {
				res[j] = res[j] + " "
			}
		}
	}
	// trim " " from the right
	for i := 0; i < maxL; i++ {
		p := -1
		for j := len(words) - 1; j >= 0; j-- {
			if res[i][j] != ' ' {
				p = j
				break
			}
		}
		res[i] = res[i][0 : p+1]
	}
	return res
}