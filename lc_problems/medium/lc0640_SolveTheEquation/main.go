func solveEquation(equation string) string {
	e := strings.Split(equation, "=")
	a, b := parse(e[0])
	c, d := parse(e[1])
	a -= c
	d -= b
	if a == 0 && d == 0 {
		return "Infinite solutions"
	}
	if a == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(d/a)
}

func parse(s string) (int, int) {
	a, b := 0, 0
	i := 0
	for i < len(s) {
		j := i
		j++
		for j < len(s) && s[j] != '+' && s[j] != '-' {
			j++
		}
		token := s[i:j]
		if s[j-1] == 'x' {
			if token == "+x" || token == "x" {
				a += 1
			} else if token == "-x" {
				a -= 1
			} else {
				temp, _ := strconv.Atoi(token[:len(token)-1])
				a += temp
			}
		} else {
			temp, _ := strconv.Atoi(token)
			b += temp
		}
		i = j
	}
	return a, b
}