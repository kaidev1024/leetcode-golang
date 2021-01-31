func interpret(command string) string {
	l := len(command)
	result := make([]byte, l)
	j := 0
	for i := 0; i < l; {
		if command[i] == 'G' {
			result[j] = 'G'
			i++
			j++
			continue
		}
		if command[i] == '(' {
			if command[i+1] == ')' {
				result[j] = 'o'
				i += 2
				j++
			} else {
				result[j] = 'a'
				result[j+1] = 'l'
				j += 2
				i += 4
			}
		}
	}
	return string(result[:j])
}