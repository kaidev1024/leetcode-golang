func fractionAddition(expression string) string {
	if len(strings.Split(expression, "/")) == 2 {
		return expression
	} else {
		i := 0
		if expression[i] == '-' || expression[i] == '+' {
			i++
		}
		for expression[i] != '+' && expression[i] != '-' {
			i++
		}
		a1 := expression[:i]
		a2 := fractionAddition(expression[i:])
		list1 := strings.Split(a1, "/")
		list2 := strings.Split(a2, "/")
		b1, _ := strconv.Atoi(list1[0])
		b2, _ := strconv.Atoi(list1[1])
		c1, _ := strconv.Atoi(list2[0])
		c2, _ := strconv.Atoi(list2[1])
		d1 := b1*c2 + b2*c1
		d2 := b2 * c2
		d1, d2 = reduce(d1, d2)
		if (d1 < 0 && d2 < 0) || (d1 > 0 && d2 < 0) {
			d1 = -d1
			d2 = -d2
		}
		return strconv.Itoa(d1) + "/" + strconv.Itoa(d2)
	}
}

func reduce(a int, b int) (int, int) {
	if a == 0 {
		return 0, 1
	}
	c := gcd(a, b)
	return a / c, b / c
}

func gcd(a int, b int) int {
	if a < 0 {
		return gcd(-a, b)
	}
	if b < 0 {
		return gcd(a, -b)
	}
	if a < b {
		return gcd(b, a)
	}
	if a == b {
		return a
	}
	return gcd(b, a-b)
}