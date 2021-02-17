func operate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0
}

func isOp(str string) bool {
	return str == "+" || str == "-" || str == "*" || str == "/"
}

func evalRPN(tokens []string) int {
	n := len(tokens)
	result := make([]int, 0, n)
	pos := 0
	for i := 0; i < n; i++ {
		if isOp(tokens[i]) {
			result[pos-2] = operate(result[pos-2], result[pos-1], tokens[i])
			pos--
			result = result[:pos]
		} else {
			v, _ := strconv.Atoi(tokens[i])
			result = append(result, v)
			pos++
		}
	}
	fmt.Println(result)
	return result[0]
}